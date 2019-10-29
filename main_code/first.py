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
    7:'change_to_daily',
    8:'change_to_normal',
}

table_state = {
    1:"events_arrangement",
    2:"daily_events_arrangement",
}
class Main_control_machine(object):
    # 初始化状态机、数据库模块
    def __init__ (self):
        self.database_control = database_mod.database_connect.Database_control()
        self.daily_table_control = database_mod.database_connect.daily_table_control()
        self.state = state_machine[1]
    # 状态机启动
    def run(self):
        temp_database = self.database_control
        temp_table = table_state[1]

        while self.state != 'end':
            if self.state == 'spare':
                temp_database.print_data(temp_table)
                try:
                    temp = int(input('input num:'))
                    self.state = state_machine[temp]
                except:
                    temp_database.print_data(temp_table)
            elif self.state == 'insert':
                self.state = temp_database.insert_into_database()
            elif self.state == 'delete':
                self.state=temp_database.delete_id_row(temp_table)
            elif self.state == 'search':
                pass
                #self.state = temp_database
            elif self.state == 'change':
                self.state = temp_database.change_database()
            elif self.state == 'change_to_daily':
                self.state = state_machine[1]
                temp_table = table_state[2]
                temp_database = self.daily_table_control
                continue
            elif self.state == 'change_to_normal':
                self.state = state_machine[1]
                temp_table = table_state[1]
                temp_database = self.database_control
                continue
            # else:
            #     temp_database.end_connect()
            #    % sys.exit(0)            
            temp_database.end_connect()

def start():
    main_control_machine = Main_control_machine()
    main_control_machine.run()

if __name__=='__main__':
    start()