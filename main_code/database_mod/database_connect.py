import pymysql
import datetime


state_machine = {
    1:'spare',
    2:'insert',
    3:'change',
    4:'delete',
    5:'end',
    6:'search',
}

class Database_control(object):
    def __init__(self):
        self.db = pymysql.connect("121.199.40.243","YuanCheng_user","_Hexagram_user","test_python")
        self.cursor = self.db.cursor()
        
    def print_data(self):
            sql = "select * from events_arrangement"
            self.__direct_database_control(sql)
            
            self.cursor.execute(sql)
            results = self.cursor.fetchall()
            for row in results:
                _id = row[0]
                _title = row[1]
                _create_date = row[2]
                _end_time = row[3]
                _advance_warning_time = row[4]
                _extre_sth = row[5]
                    
                print("%s,%s,%s,%s,%s,%s" % \
                    (_id,_title,_create_date, \
                    _end_time,_advance_warning_time,_extre_sth))

    def insert_into_database(self):
        title = input("please input sth")
        end_time = input("please input end time,\"- -\"")
        if end_time == '':
            end_time = '1111-11-11 11:11:11'

        advance_warning_time = input("please input advance warning tima,\": :\"")
        if advance_warning_time == '':
            advance_warning_time = '11:11:11'
            
        extre_sth = input("please inpot extra sth")
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
        self.__direct_database_control(sql)
        return state_machine[1]

    def delete_id_row(self):
        row_id = int(input("Which id do you want to delete"))
        sql = """
            DELETE FROM events_arrangement
            WHERE id = %d
            """ % (row_id)
        self.__direct_database_control(sql)
        return state_machine[1]

    def __direct_database_control(self,sql):
        try:
            self.cursor.execute(sql)
            self.db.commit()
        except:
            print("wrong!!!")
            self.db.rollback()

    def end_connect(self):
        self.db.close()