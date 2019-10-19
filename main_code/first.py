#!/usr/bin/env python 
# -*- coding: utf-8 -*- j
import sys
import database_mod.database_connect

state_machine = {
    1:'spare', #
    2:'insert',#
    3:'change',#
    4:'delete',#
    5:'end',#
    6:'search',
}
class Main_control_machine(object):
    def __init__ (self):
        self.database_control = database_mod.database_connect.Database_control()
        self.state = state_machine[1]
    def run(self):
        while self.state != 'end':
            if self.state == 'spare':
                self.database_control.print_data()
                self.state = state_machine[int(input('input num'))]
            if self.state == 'insert':
                self.state = self.database_control.insert_into_database()
            if self.state == 'delete':
                self.state=self.database_control.delete_id_row()
            if self.state == 'search':
                pass
                #self.state = self.database_control
            if self.state == 'change':
                self.state = self.database_control.change_database()
        else:
            self.database_control.end_connect()
            sys.exit(0)            

def start():
    main_control_machine = Main_control_machine()
    main_control_machine.run()

if __name__=='__main__':
    print(state_machine)
    start()