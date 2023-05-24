def calc_tri(base, height):
    return base*height * (1/2)

base = 10
heght = 'Hello'

# try:
#     エラーの起こりうる処理
# except Exception as e: <-どのエラーにも対応
#     エラー発生時の処理
# finally:
#     必ず処理させたい処理

try:
    result =calc_tri(base, heght)
    print('計算結果:{}'.format(result))
except Exception as e:
    print('計算できません')
    print(e) #エラー内容の表示

try:
    a = 10
    b = 1
    result = a/b
except ZeroDivisionError:
    print('0で割っています。')
finally:
    print("必ず通る。")