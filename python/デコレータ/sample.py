#関数の引数に関数
def func1():
    print('Called func1')

def func2(f):
    f()

print(func1) #<-関数は、objectである。
print(func2(func1))



#関数内関数 パターン1
#func3は、func4に対してデコレータとして働いている。
def func3(f):
    def wrapper():
        print('開始')
        f()
        print('終了')
    return wrapper #<-wapperobjecを返している。

def func4():
    print('Called func4')

func = func3(func4) #<-funcには、wapper objectが代入されている。
func() # <- wapper()と同義
func3(func4)() # = func()


#関数内関数 パターン2
import datetime
def print_datetime(f):
    def wrapper():
        print(f'開始: {datetime.datetime.now()}')
        f()
        print(f'終了: {datetime.datetime.now()}')
    return wrapper

def calc():
    print(3**5)

print_datetime(calc)()


#関数内関数 デコレータの利用
import datetime
def print_datetime(f):
    def wrapper():
        print(f'開始: {datetime.datetime.now()}')
        f()
        print(f'終了: {datetime.datetime.now()}')
    return wrapper

@print_datetime #<- デコレータの追加
def calc():
    print(3**5)

calc()


#引数を受け取る関数デコレータ @を利用しない場合
import datetime
def print_datetime(f):
    def wrapper(base, height):
        print(f'開始: {datetime.datetime.now()}')
        f(base, height)
        print(f'終了: {datetime.datetime.now()}')
    return wrapper

def calc(base, height):
    print(base * height / 2)

print_datetime(calc)(3, 10)
#wapper = print_datetime(calc)
#wapper(3, 10) =  print_datetime(calc)(3, 10)


#引数を受け取る関数デコレータ デコレータを利用する場合
import datetime
def print_datetime(f):
    def wrapper(base, height):
        print(f'開始: {datetime.datetime.now()}')
        f(base, height)
        print(f'終了: {datetime.datetime.now()}')
    return wrapper

@print_datetime
def calc(base, height):
    print(base * height / 2)

calc(3, 10)


#可変長引数を用いたデコレータ
import datetime
def print_datetime(f):
    def wrapper(*args, **kwargs):
        print(f'開始: {datetime.datetime.now()}')
        f(*args, **kwargs)
        print(f'終了: {datetime.datetime.now()}')
    return wrapper

@print_datetime
def calc1(base, height):
    print(base * height / 2)

@print_datetime
def calc2(a, b, c):
    print(a*b*c)

calc1(3, 10)
calc2(3, 10, 6)