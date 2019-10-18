#!/usr/bin/env python 
# -*- coding: utf-8 -*- 
import sys
import pymysql
import datetime

state_machine = {
    1:'spare',
    2:'insert',
    3:'',
    4:'delete',
    5:'end',
    6:'search',
}
class Main_control_machine(object):
    def __init__ (self):
        self.database_control = Database_control()
        self.state = state_machine[1]
    def run(self):
        while self.state != 'end':
            if self.state == 'spare':
                self.database_control.print_data()
                self.state = state_machine[int(input('input num'))]
            if self.state == 'insert':
                self.state = self.database_control.insert_into_database()
            if self.state == 'delete':
                pass
            if self.state == 'search':
                pass
        else:
            self.database_control.end_connect()
            sys.exit(0)            

class Database_control(object):
    def __init__(self):
        self.db = pymysql.connect("121.199.40.243","YuanCheng_user","_Hexagram_user","test_python")
        self.cursor = self.db.cursor()
    def print_data(self):
            sql = "select * from events_arrangement"
        
            try:
                self.cursor.execute(sql)
                results = self.cursor.fetchall()
                for row in results:
                    pass
            except:
                print("error in fetch data")
            print(results)

    def insert_into_database(self):
        title = input("please input sth")
        end_time = input("please input end time,\"- -\"")
        advance_warning_time = input("please input advance warning tima,\": :\"")
        extre_sth = input("please inpot extra sth")
        create_time = datetime.datetime.now()
        create_time_str = datetime.datetime.strftime(create_time,'%Y-%m-%d %H:%M:%S')

        sql = """
            insert into table events_arrangement 
            (events_title,events_create_date,events_end_time,events_advance_waring_time,extra_sth)
            values
            (%s,%s,%s,%s,%s)
            """ % (title,create_time_str,end_time,advance_warning_time,extre_sth)
        #try:
        self.cursor.execute(sql)
        self.db.commit()
        #except:
        #    print("wrong!!!")
        #    self.db.rollback()
        return state_machine[1]
    def end_connect(self):
        self.db.close()

def start():
    main_control_machine = Main_control_machine()
    main_control_machine.run()

if __name__=='__main__':
    print(state_machine)
    start()