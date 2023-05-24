#何個でも引数を代入できる
def func(*args):
    print(args)

func(1)
func(1, 3, 10)

def func_join(*args2):
    result = ','.join(args2)
    print(result)

func_join('あ')
func_join('あ', 'い', 'う')

#*argsとただの変数を同時に利用する場合
##普通の引数が先x
def funcX(x, *args):
    print(args)
    print(x)
funcX(1, 'a', 'b', 'c')

##普通の引数が後
def funcX(*args, x):
    print(args)
    print(x)
funcX('a', 'b', 'c', x=1)


#辞書式で複数の可変の引数を設定する
def func3(**kwargs):
    print(kwargs)

func3(name='斎藤')
func3(name='鈴木', id=1, type='02')