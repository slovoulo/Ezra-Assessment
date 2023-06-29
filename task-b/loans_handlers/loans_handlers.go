// @title Loans App API
// @version 1.0.0
// @description Loans API documentation.
// @host localhost:7071
// @Accept json
// @Produce json
// @BasePath /
package loans_handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/slovoulo/Ezra-Assessment/task-b/loans_database"
	"github.com/slovoulo/Ezra-Assessment/task-b/loans_models"
	// "golang.org/x/crypto/ssh"
	// "github.com/pkg/sftp"
)

type SwaggerAccID struct{
	AccountID int
}
type SwaggerLoanReq struct{
	AccountID int
	AmountAdded int

}

// Home godoc
// @Summary Loans Landing page
// @Description Landing page
// @Tags Loans-App
// @Success 200 
// @Failure 400 
// @Router /v1/ [get]
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Loans app!"))

}

func checkMSIDN(w http.ResponseWriter,value int) error{
	//This flow assumes MSIDN has 10 characters
	 //IS an integer and not negative

	 // Check if the value is negative
	 if value < 0 {
		log.Println("Account ID/MSIDN cannot be negative")
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("Account ID/MSIDN cannot be negative"))
        return errors.New("Account ID/MSIDN cannot be negative")
    }
    
    // Convert the integer to a string
    strValue := strconv.Itoa(value)
    
    // Check the length of the string
    if len(strValue) != 9 {
		log.Println("Account ID/MSIDN must have exactly 9 characters")
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("Account ID/MSIDN must have exactly 9 characters"))
        return  errors.New("Account ID/MSIDN must have exactly 9 characters")
    }
    // Make sure msdin doesnt start with 0
    if strValue[0] == 0 {
		log.Println("Account ID/MSIDN cannot start with a 0")
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("Account ID/MSIDN cannot start with a 0"))
        return  errors.New("Account ID/MSIDN cannot start with a 0")
    }
	return nil
    
    
}


///Annotation for creating a loan account

// @Summary Create loan account
// @Description Create loan account
// @Description MUST NOT start with a 0
// @Tags Loans-App
// @Accept  json
// @Produce  json
// @Success 200 
// @Failure 400 
// @Param AccountID body SwaggerAccID true "Loan account ID (MSDIN)"
// @Router /v1/account [post]
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	//Before requesting a loan, user should have an account
	// get the body of the  POST request
	// unmarshal this into a new User struct
	// append this to the Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account loans_models.Account
	 json.Unmarshal(reqBody, &account)

	 
	 //Check if MSIDN is valid
	 msderr:=checkMSIDN(w,account.AccountID)
	 if msderr!=nil{
		
		return
	 }
	 //Set new users loan limit to 2000
	 account.LoanLimit=2000
	 //Set created time to time.now
	 account.Created_at=time.Now()

	

	if result := loans_database.Db.Create(&account); result.Error != nil {
		fmt.Println(result.Error)
	}
	
	json.NewEncoder(w).Encode(account)
	log.Println("Account created successfully")
	w.WriteHeader(http.StatusOK)
    w.Write([]byte("Account cretaed successfully"))

}

//Check user's account in the database to see their loan limit and unpaid balance
func GetUserLoanDetails(w http.ResponseWriter, accountID int) (int,int,error) {
	
	var userAccount loans_models.Account
	//SELECT * FROM databaseTable WHERE aaccountid==aaccountID
	res := loans_database.Db.Find(&userAccount, "account_id = ?", accountID)
	if res.Error!=nil{
		errorMessage:=fmt.Sprintf("There was an error fetching user records %s", res.Error)
		log.Printf(errorMessage)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errorMessage))
		return 0,0, errors.New(errorMessage)

	}
	// Marshal the record into a JSON string
     jsonStr, err := json.Marshal(userAccount)
	 if(err!=nil){
		log.Println("An error occcured unmarshaling json")
		return 0,0, errors.New("An error occcured unmarshaling json")
	 }

	 // Unmarshal the record into a userAccount struct
     var userAc loans_models.Account
    err1 := json.Unmarshal([]byte(jsonStr), &userAc)
	if err1 !=nil{
		log.Printf("An error occured Unmarshaling json : %s", err1)
		return 0,0,errors.New("An error occcured unmarshaling json")
	}

	return userAc.LoanLimit, userAc.Unpaid_balance,nil
	

}


///Annotation for requesting a loan

// @Summary Request a loan
// @Description Request loan
// @Description "Amountadded" is the loan amount you are requesting
// @Tags Loans-App
// @Accept  json
// @Produce  json
// @Success 200 
// @Failure 400 
// @Param LoanDetails body SwaggerLoanReq true "Loan account ID (MSDIN) and Amount"
// @Router /v1/loanrequest [post]
func LoanRequest(w http.ResponseWriter, r *http.Request){

	// get the body of the  POST request
	// unmarshal this into a new User struct
	// append this to the Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var transaction loans_models.TransactionEntries
	 json.Unmarshal(reqBody, &transaction)

	 

	//Check users loan limit
	loanLimit, unpaidBal, detailsErr:=GetUserLoanDetails(w, transaction.AccountID)
	//Compare loan limit to requested amount
	if detailsErr!=nil{
		//If user is requesting above their limit
		msg:=fmt.Sprintf("An error occurred %s", detailsErr)
		log.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}else if loanLimit< transaction.AmountAdded{
		//If user is requesting above their limit
		log.Printf("Error: You cannot request more than your loan limit")
		w.WriteHeader(http.StatusBadRequest)
		mess:=fmt.Sprintf("Error: You cannot request more than your loan limit of %s", strconv.Itoa(loanLimit))
		w.Write([]byte(mess))
		return
		} else if transaction.AmountAdded<10{
			//If user is requesting above their limit
		log.Printf("Error: You cannot request less than 10/=")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: You cannot request less than 10/="))
		return
		}else if loanLimit >= transaction.AmountAdded{
			///If loan limit is okay, perform request and then:
			//1. Record the transaction
			transaction.AmountAdded=-(transaction.AmountAdded) //Since its a borrow, convert the int to negative
			transaction.TransactionType="Loan Request"
			if result := loans_database.Db.Create(&transaction); result.Error != nil {
				mess:=fmt.Sprintf("An error occured recording the transaction %s", result.Error)
				log.Println(mess)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(mess))
				return 
			}
				log.Printf("Transaction recorded successfully")
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("Transaction recorded successfully\n"))

			

			//2.Update their unpaid balance
			newUnpaid:=unpaidBal - transaction.AmountAdded //Remember we converted transaction.AmountAdded to negative


			//3.Update their loan limit
			newLoanLimit:=loanLimit+transaction.AmountAdded //Subtract the approved amount from the existing loan limit

			newlyUpdatedAccount:=loans_models.Account{
				Unpaid_balance: newUnpaid,
				LoanLimit: newLoanLimit,
				
			}
			
			//Perform update on DB
			//Update recipe
		result := loans_database.Db.Where("account_id = ?", transaction.AccountID).Updates(newlyUpdatedAccount)
		if result.Error != nil {
		log.Printf("There was an error updating Account %s ", result.Error)
		w.WriteHeader(http.StatusNotFound)
		return

	}
	json.NewEncoder(w).Encode(newlyUpdatedAccount)
	log.Printf("Account updated successfully/n")
	w.Write([]byte("Successfully updated account : "))



		}

	


}
///Annotation for repaying a loan

// @Summary Repay a loan
// @Description Repay loan
// @Tags Loans-App
// @Accept  json
// @Produce  json
// @Success 200 
// @Failure 400 
// @Param LoanDetails body SwaggerLoanReq true "Loan account ID (MSDIN) and Amount"
// @Router /v1/loanrepayment [post]
func LoanRepayment(w http.ResponseWriter, r *http.Request){

	// get the body of the  POST request
	// unmarshal this into a new User struct
	// append this to the Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var transaction loans_models.TransactionEntries
	 json.Unmarshal(reqBody, &transaction)

	 

	 loanLimit, unpaidBal, detailsErr:=GetUserLoanDetails(w, transaction.AccountID)
	//Compare loan limit to requested amount
	if detailsErr!=nil{
		//If user is requesting above their limit
		msg:=fmt.Sprintf("An error occurred %s", detailsErr)
		log.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}else if unpaidBal< transaction.AmountAdded{
		//If user wants to pay more than they owe
		log.Printf("Error: You cannot pay more than you owe")
		w.WriteHeader(http.StatusBadRequest)
		mess:=fmt.Sprintf("Error: You cannot pay more than your loan balance  of %s", strconv.Itoa(unpaidBal))
		w.Write([]byte(mess))
		return
		} else if unpaidBal >= transaction.AmountAdded{
			///If loan limit is okay, perform request and then:
			//1. Record the transaction
			
			transaction.TransactionType="Loan Repayment"
			if result := loans_database.Db.Create(&transaction); result.Error != nil {
				mess:=fmt.Sprintf("An error occured recording the transaction %s", result.Error)
				log.Println(mess)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(mess))
				return 
			}
				log.Printf("Transaction recorded successfully")
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("Transaction recorded successfully/n"))

			

			//2.Update their unpaid balance
			newUnpaid:=unpaidBal - transaction.AmountAdded 


			//3.Update their loan limit
			newLoanLimit:=loanLimit+transaction.AmountAdded 

			newlyUpdatedAccount:=loans_models.Account{
				Unpaid_balance: newUnpaid,
				LoanLimit: newLoanLimit,
				
			}
			
			//Perform update on DB
			//Update recipe
		result := loans_database.Db.Where("account_id = ?", transaction.AccountID).Updates(newlyUpdatedAccount)
		if result.Error != nil {
		log.Printf("There was an error updating Account %s ", result.Error)
		w.WriteHeader(http.StatusNotFound)
		return

	}
	json.NewEncoder(w).Encode(newlyUpdatedAccount)
	log.Printf("Account updated successfully/n")
	w.Write([]byte("Successfully updated account : "))



		}

	


}


//Generate dumps

func generateDump(loans []loans_models.TransactionEntries) {
    // Create a new file to store the dump
    dumpFile, err := os.Create("loans.dump")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Write the loans to the file
    data, err := json.Marshal(loans)
    if err != nil {
        fmt.Println(err)
        return
    }

    dumpFile.Write(data)
    dumpFile.Close()
}

// SFTP configuration
const (
	SFTPServer   = "EZRA_SFTP_SERVER"
	SFTPUsername = "EZRA_SFTP_USERNAME"
	SFTPPassword = "EZRA_SFTP_PASSWORD"
	SFTPPath     = "EZRA_SFTP_PATH"
)

// // Upload dump file to SFTP server
// func uploadDumpToSFTP(dumpFileName string) error {
// 	// Read the dump file
// 	dumpData, err := ioutil.ReadFile(dumpFileName)
// 	if err != nil {
// 		return fmt.Errorf("failed to read the dump file: %s", err.Error())
// 	}

// 	// Create SFTP client configuration
// 	config := &ssh.ClientConfig{
// 		User: SFTPUsername,
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password(SFTPPassword),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}

// 	// Connect to the SFTP server
// 	sftpClient, err := sftp.NewClient(SFTPServer, config)
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to SFTP server: %s", err.Error())
// 	}
// 	defer sftpClient.Close()

// 	// Create the remote file on the SFTP server
// 	remoteFile, err := sftpClient.Create(SFTPPath + "/" + dumpFileName)
// 	if err != nil {
// 		return fmt.Errorf("failed to create remote file on SFTP server: %s", err.Error())
// 	}
// 	defer remoteFile.Close()

// 	// Write the dump data to the remote file
// 	_, err = remoteFile.Write(dumpData)
// 	if err != nil {
// 		return fmt.Errorf("failed to write data to remote file: %s", err.Error())
// 	}
// 	return nil
// }