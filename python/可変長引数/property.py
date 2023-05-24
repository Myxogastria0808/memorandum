#デコレータなし
class  User:
    def __init__(self):
        self._name = None 
    #nameを取得する
    def get_name(self): #<-getter
        return self._name

    def set_name(self, name): #<-setter
        if type(name) == str:
            self._name = name
        else:
            raise TypeError('文字列にしてください')

user1 = User()
#ただのget_name()
print(user1.get_name())
#デコレータなしの場合の使い方
user1.set_name('Suzuki')
print(user1.get_name())


#デコレータあり
class  UserDeco:
    def __init__(self):
        self._name = None

    #nameを取得する
    @property 
    def name(self): #<-getter
        return self._name

    @name.setter
    def name(self, name): #<-setter
        if type(name) == str:
            self._name = name
        else:
            raise TypeError('文字列にしてください')

user1 = UserDeco()
user1.name = 'Suzuki'
print(user1.name)