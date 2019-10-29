#!/usr/bin/env python 
# -*- coding: utf-8 -*- j
import pymysql
import datetime
import prettytable
import os
import time
import sys

def log(func):
    def wrapper(*args,**kw):
        print('call %s():' % (func.__name__))
        return func(*args,**kw)
    return wrapper

state_machine = {
    1:'spare',
    2:'insert',
    3:'change',
    4:'delete',
    5:'end',
    6:'search',
    7:'change_to_daily',
    8:'change_to_normal',
}

class Database_control(object):
    def __init__(self):
        pass

    def print_data(self,temp_table):
        os.system('cls')
        print(state_machine)
        create_time = datetime.datetime.now()
        create_time = datetime.datetime.strftime(create_time,'%Y-%m-%d %H:%M:%S')
        print(create_time)
        print()
        sql = "select * from %s" % (temp_table)
        self._direct_database_control(sql)
            
        # self.cursor.execute(sql)
        results = self.cursor.fetchall()
        print(self._direct_print(results))

    @log
    def change_database(self):
        colums = {
            1:'events_end_time',
            2:'events_title',
            3:'events_advance_waring_time',
            4:'extre_sth'
        }
        temp_id = int(input("please the id you want to change:"))
        temp=int(input("""please what you want to change\n 
            (end_time:1, sth_title:2, advance_warning_time:3,extre_sth:4):"""))
        
        if temp == 1:
            temp_input = input("end_time:\n")
        elif temp == 2:
            temp_input = input("sth_title:\n")
        elif temp == 4:
            temp_input = input("extre_sth\n")
        else:
            # temp == 3
            temp_input = input("advance_warning_time:\n")
        sql = \
            """
            UPDATE events_arrangement
            SET %s='%s'
            WHERE id=%d  
            """ % (colums[temp],temp_input,temp_id)
        self._direct_database_control(sql)
        
        return state_machine[1]

    @log
    def insert_into_database(self):
        title = input("please input sth:\n")
        end_time = input("please input end time,\"- -\":"  )
        if end_time == '':
            end_time = '1111-11-11 00:00:00'

        advance_warning_time = input("please input advance warning tima,\": :\":")
        if advance_warning_time == '':
            advance_warning_time = '00:00:00'
            
        extre_sth = input("please inpot extra sth:\n")
        create_time = datetime.datetime.now()
        create_time_str = datetime.datetime.strftime(create_time,'%Y-%m-%d %H:%M:%S')

        sql = """
            insert into events_arrangement 
            (events_title,events_create_date,events_end_time,
            events_advance_waring_time,extre_sth)
            values
            ('%s','%s','%s','%s','%s')
            """ % (title,create_time_str,end_time,advance_warning_time,extre_sth)
            #print(sql)
        self._direct_database_control(sql)
        return state_machine[1]

    @log
    def delete_id_row(self):
        row_id = int(input("Which id do you want to delete:"))
        sql = """
            SELECT * FROM events_arrangement
            WHERE id = %d
            """ % (row_id)
        self._direct_database_control(sql)
        results = self.cursor.fetchall()
        self._insert_into_file(results)

        sql = """
            DELETE FROM events_arrangement
            WHERE id = %d
            """ % (row_id)
        self._direct_database_control(sql)
        return state_machine[1]

    def _direct_database_control(self,sql):
        self._connect_database()
        try:
            self.cursor.execute(sql)
            self.db.commit()
        except:
            print("wrong!!!")
            self.db.rollback()
            time.sleep(3)

    def _direct_print(self,temp):
        table = prettytable.PrettyTable(["ID","事务内容", \
            "创建时间","截至日期","提前预警时间","额外附加"])
        for row in temp:
            _id = row[0]
            _title = row[1]
            _create_date = str(row[2])
            _end_time = str(row[3])
            _advance_warning_time = str(row[4])
            _extre_sth = row[5]            
            table.add_row( \
            [_id,_title,_create_date[0:-3], \
            _end_time[0:-3],_advance_warning_time[0:-3],_extre_sth] \
            )
        return table

    def _insert_into_file(self,file_neirong):
        sth_to_say = input("please say sth about this events:\n")
        path1 = os.path.abspath('.')
        path1 = path1 + '/main_code/database_log/log.txt'
        # print(path1)
        temp = self._direct_print(file_neirong)
        f = open(path1,mode='a',encoding='utf-8')
        f.write("\n\n%s\n%s" % (temp,sth_to_say))
        f.close

    def _connect_database(self):
        try:
            self.db = pymysql.connect("121.199.40.243","YuanCheng_user","_Hexagram_user","test_python")
            self.cursor = self.db.cursor()
        except:
            time.sleep(3)
            self._connect_database()

    def end_connect(self):
        self.db.close()

class daily_table_control(Database_control):
    def _direct_print(self,temp):
        table = prettytable.PrettyTable(["ID","事务内容", \
            "推荐时间点","额外附加"])
        for row in temp:
            _id = row[0]
            _title = row[1]
            do_it_time = str(row[2])
            _extre_sth = row[3]            
            table.add_row( \
            [_id,_title,do_it_time, \
            _extre_sth] \
            )
        return table

    @log
    def insert_into_database(self):
        title = input("please input sth:\n")
        events_to_do_time = input("please input events_to_do_time,\": :\":"  )
        if events_to_do_time == '':
            events_to_do_time = '00:00:00'

        extre_sth = input("please inpot extra sth:\n")

        sql = """
            insert into daily_events_arrangement 
            (events_title,events_do_it_time,extre_sth)
            values
            ('%s','%s','%s')
            """ % (title,events_to_do_time,extre_sth)
            #print(sql)
        self._direct_database_control(sql)
        return state_machine[1]