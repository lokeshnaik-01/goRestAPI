# restAPI

db.DB.Query is used when we want to get some rows

db.DB.Exec is used when it changes db like write or update

And we did this by following different approaches:

1. DB.Exec() (when we created the tables)

2. Prepare() + stmt.Exec() (when we inserted data into the database)

3. DB.Query() (when we fetched data from the database)

Using Prepare() is 100% optional! You could send all your commands directly via Exec() or Query().

The difference between those two methods then just is whether you're fetching data from the database (=> use Query()) or your manipulating the database / data in the database (=> use Exec()).

But what's the advantage of using Prepare()?

Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).

This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.

And, indeed, in this application, we are calling stmt.Close() directly after calling stmt.Exec(). So here, it really wouldn't matter which approach you're using.