DB Archiver is a utility developed in golang useful for archiving 
data. Currently this utility only supports archiving data to CSV 
file. 

##### Installation
you can download the executable(LINUX only) from github or 
you can clone the repository to your GOPATH. You need glide if
you are going to build from source

<pre>
<code>
1. git clone git@github.com:traveltriangle/db-archiver.git
2. glide install
3. go build -o db-archiver app.go 
</code>
</pre>

you can check for an executable with the name db-archiver.

##### How to use
In the installation path 
<pre>
<code>
1. Create a folder config
2. Inside config create a file db.yaml and add the 
   following configuration for your database
    <b>user</b>: admin
    <b>password</b>: password
    <b>address</b>: localhost:3306

</code>
</pre>
you can now use db-archiver to archive any table to a CSV file which can
then be moved to any storage space like S3 or Glacier.

##### Examples



<pre><code>
./db-archiver --db-name <b>DB_NAME</b> \
--table <b>TABLE_NAME</b> --where <b>CONDITION</b>

./db-archiver --db-name <b>DB_NAME</b> \
--table <b>TABLE_NAME</b> --query <b>QUERY</b>

./db-archiver --db-name <b>DB_NAME</b> \
--table <b>TABLE_NAME</b> --where <b>CONDITION</b> \
--path <b>FOLDER_PATH</b>

./db-archiver --db-name test \
--table customers --where "join_date > '2019-01-01'" \
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
   <b>--path </b>string
     	path to folder where the file will be stored (default "/tmp/")
   <b>--pk </b>string
     	primary key which will be used to delete the records (default "id")
   <b>--query </b>string
     	if used it will ignore <b>--where </b>option. Needed if <b>--where </b>is not provided
   <b>--table </b>string
     	[REQUIRED] table name to be archived
   <b>--where </b>string
     	condition to be used while archiving. Needed if <b>--query </b>is not provided
</code></pre>
