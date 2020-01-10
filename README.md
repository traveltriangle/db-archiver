DB Archiver is a utility developed in golang useful for archiving 
data. Currently this utility only supports archiving data to CSV 
file. 

##### Installation 

Clone the repository to your GOPATH and build the app.go

<pre>
<code>
git clone git@github.com:traveltriangle/db-archiver.git
go build -o db-archiver app.go
</code>
</pre>

you can check for an executable with the name db-archiver.
This executable can be moved to any folder in the PATH variable.

##### How to use

We can archive any table with this utility to a CSV file which can
then be moved to any storage space like S3 or Glacier.

###### Examples



<pre><code>
./db-archiver --db-name <b>DB_NAME</b> \
--table <b>TABLE_NAME</b> --where <b>CONDITION</b>

./db-archiver --db-name <b>DB_NAME</b> \
--table <b>TABLE_NAME</b> --query <b>QUERY</b>

./db-archiver --db-name <b>DB_NAME</b> \
--table <b>TABLE_NAME</b> --where <b>CONDITION</b> \
--user <b>USER</b> --password <b>PASSWORD</b> \
--path <b>FOLDER_PATH</b>

./db-archiver --db-name test \
--table customers --where "join_date > '2019-01-01'" \
--user admin --password password \
--path "/var/log/"

./db-archiver --db-name test --table customers \
--query "select customers.id as id, customers.name as name, join_date 
  from customers 
  inner join depts on
  depts.customer_id = customers.id
  where depts.name = 'development'" \
--user admin --password password \
--path "/var/log/"




<b><i>Available Options</i></b>
   <b>--address</b> string
     	[REQUIRED] Host address with port of Database Connection. (default "localhost:3306")
   <b>--batch</b> int
     	Fetch records in batch. If used it will ignore <b>--limit</b>
   <b>--db-name</b> string
     	[REQUIRED] Name of Database
   <b>--delete</b>
     	Delete from Table after archiving (default true)
   <b>--limit </b>int
     	limit the number of records (default 500)
   <b>--optimize</b>
     	Optimize Table after deletion (default true)
   <b>--password </b>string
     	Password for Database Connection
   <b>--path </b>string
     	path to folder where the file will be stored (default "/tmp/")
   <b>--pk </b>string
     	primary key which will be used to delete the records (default "id")
   <b>--query </b>string
     	if used it will ignore <b>--where </b>option. Needed if <b>--where </b>is not provided
   <b>--table </b>string
     	[REQUIRED] table name to be archived
   <b>--user </b>string
     	User for Database Connection
   <b>--where </b>string
     	condition to be used while archiving. Needed if <b>--query </b>is not provided
</code></pre>
