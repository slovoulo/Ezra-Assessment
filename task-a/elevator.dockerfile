#build a tiny docker image
FROM alpine:latest

RUN mkdir /app


COPY elevatorapp /app



RUN chmod +x /app/elevatorapp




#Build the first docker image, create a  much smaller docker image then copy the executable from first to second smaller image
CMD [ "/app/elevatorapp" ]