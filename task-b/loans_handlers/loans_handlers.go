package loans_handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/slovoulo/Ezra-Assessment/task-b/loans_database"
	"github.com/slovoulo/Ezra-Assessment/task-b/loans_models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Loans app!"))

}

func checkMSIDN(w http.ResponseWriter,value int){
	//This flow assumes MSIDN has 10 characters
	 //IS an integer and not negative

	 // Check if the value is negative
	 if value < 0 {
		log.Println("MSIDN cannot be negative")
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("MSDIN cannot be negative"))
        return 
    }
    
    // Convert the integer to a string
    strValue := strconv.Itoa(value)
    
    // Check the length of the string
    if len(strValue) != 9 {
		log.Println("MSIDN must have 9 characters")
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("MSDIN must have 9 characters"))
        return 
    }
    // Make sure msdin doesnt start with 0
    if strValue[0] == 0 {
		log.Println("MSIDN cannot start with a 0")
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("MSIDN cannot start with a 0"))
        return 
    }
    
    
}

//Before requesting a loan, user should have an account
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	// get the body of the  POST request
	// unmarshal this into a new User struct
	// append this to the Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account loans_models.Account
	 json.Unmarshal(reqBody, &account)

	 
	 //Check if MSIDN is valid
	 checkMSIDN(w,account.AccountID)
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
func GetUserLoanLimit(w http.ResponseWriter, accountID int) (int,int) {
	
	var userAccount loans_models.Account
	//SELECT * FROM databaseTable WHERE aaccountid==aaccountID
	res := loans_database.Db.Find(&userAccount, "accountid = ?", accountID)
	if res.Error!=nil{
		log.Printf("There was an error fetching user records %s", res.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("There was an error fetching user records"))
		return 0,0

	}
	// Marshal the record into a JSON string
     jsonStr, err := json.Marshal(userAccount)
	 if(err!=nil){
		log.Println("An error occcured unmarshaling json")
		return 0,0
	 }

	 // Unmarshal the record into a userAccount struct
     var userAc loans_models.Account
    err1 := json.Unmarshal([]byte(jsonStr), &userAc)
	if err1 !=nil{
		log.Printf("An error occured Unmarshaling json : %s", err1)
		return 0,0
	}

	return userAc.LoanLimit, userAc.Unpaid_balance
	

}

//Request loan
func LoanRequest(w http.ResponseWriter, r *http.Request){

	// get the body of the  POST request
	// unmarshal this into a new User struct
	// append this to the Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var transaction loans_models.TransactionEntries
	 json.Unmarshal(reqBody, &transaction)

	//Check users loan limit
	loanLimit, unpaidBal:=GetUserLoanLimit(w, transaction.AccountID)
	//Compare loan limit to requested amount
	if loanLimit< transaction.AmountAdded{
		//If user is requesting above their limit
		log.Printf("Error: You cannot request more than your loan limit")
		w.WriteHeader(http.StatusBadRequest)
		mess:=fmt.Sprintf("Error: You cannot request more than your loan limit of %s", strconv.Itoa(loanLimit))
		w.Write([]byte(mess))
		return
		} else if loanLimit >= transaction.AmountAdded{
			///If loan limit is okay, perform request and then:
			//1. Record the transaction
			transaction.AmountAdded=-(transaction.AmountAdded) //Since its a borrow, convert the int to negative
			if result := loans_database.Db.Create(&transaction); result.Error != nil {
				mess:=fmt.Sprintf("An error occured recording the transaction %s", result.Error)
				log.Printf(mess)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(mess))
				return 
			}
				log.Printf("Transaction recorded successfully")
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("Transaction recorded successfully"))

			

			//2.Update their unpaid balance
			newUnpaid:=unpaidBal - transaction.AmountAdded //Remember we converted transaction.AmountAdded to negative


			//3.Update their loan limit
			newLoanLimit:=loanLimit-transaction.AmountAdded //Subtract the approved amount from the existing loan limit
			



		}

	


}