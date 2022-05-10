from flask import Flask, render_template, request
import psycopg2
import os

app = Flask(__name__)

@app.route('/')
def checkDB():

    #This method tries to connect to the default Postgres database and gets a list of databases - can be used to verify that the database exists
    conn = None

    db_host = os.getenv("POSTGRES_HOST")
    db_name = os.getenv("POSTGRES_DB")
    db_user = os.getenv("POSTGRES_USER")
    db_password = os.getenv("POSTGRES_PASSWORD")
    db_port = os.getenv("POSTGRES_PORT")

    message_to_display = ""

    try:
        conn = psycopg2.connect(database=db_name, user=db_user, password=db_password, host=db_host, port= db_port)
    except Exception as e:
        message_to_display = "<h1>There was a problem connecting to the database! </h1>" + str(e)
    if conn is not None:
        conn.autocommit = True
        cur = conn.cursor()
        cur.execute("SELECT datname FROM pg_database;")
        database_list = cur.fetchall()
        message_to_display += "<h1>The following databases are on this postgres instance: </h1><br>"
        for database in database_list:
            message_to_display += str(database) + "<br>"        
        cur.close()
        conn.close()    
    return message_to_display
if __name__ == '__main__':
    app.run(host='0.0.0.0')