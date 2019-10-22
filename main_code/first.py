#!/usr/bin/env python 
# -*- coding: utf-8 -*- j
import sys
import database_mod.database_connect

# 状态机 #代表已完成功能
state_machine = {
    1:'spare', #
    2:'insert',#
    3:'change',#
    4:'delete',#
    5:'end',#
    6:'search',
}
class Main_control_machine(object):
    # 初始化状态机、数据库模块
    def __init__ (self):
        self.database_control = database_mod.database_connect.Database_control()
        self.state = state_machine[1]
    # 状态机启动
    def run(self):
        while self.state != 'end':
            if self.state == 'spare':
                self.database_control.print_data()
                try:
                    temp = int(input('input num:'))
                    self.state = state_machine[temp]
                except:
                    self.database_control.print_data()
            elif self.state == 'insert':
                self.state = self.database_control.insert_into_database()
            elif self.state == 'delete':
                self.state=self.database_control.delete_id_row()
            elif self.state == 'search':
                pass
                #self.state = self.database_control
            elif self.state == 'change':
                self.state = self.database_control.change_database()
            else:
                self.database_control.end_connect()
                sys.exit(0)            
            self.database_control.end_connect()

def start():
    main_control_machine = Main_control_machine()
    main_control_machine.run()

if __name__=='__main__':
    start()