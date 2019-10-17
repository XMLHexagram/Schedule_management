#!/usr/bin/env python 
# -*- coding: utf-8 -*- 
import sys
import pymysql

state_machine = {
    1:'spare',
    2:'',
    3:'input',
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
            if self.state == 'input':
                self.state = self.database_control.input_into_database(input())                  
            if self.state == 'delete':
                pass
            if self.state == 'search':
                pass
        else:
            sys.exit(0)            

class Database_control(object):
    def __init__(self):
        self.db = pymysql.connect("121.199.40.243","YuanCheng_user","_Hexagram_user","test_python")
        cursor = self.db.cursor()
        cursor.execute("desc events_arrangement")
        data = cursor.fetchall()
        print(data)
    def print_data(self):
        print(self.db)
    def input_into_database(self,input_sth):
        pass
        return state_machine[1]
    def end_connect(self):
        self.db.close()

def start():
    main_control_machine = Main_control_machine()
    main_control_machine.run()

if __name__=='__main__':
    print(state_machine)
    start()