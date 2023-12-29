from build.go_bind import TEstData, HelloThere


ex = TEstData(Data="test", Len=0)
pp = HelloThere(ex)

print(f" >> {ex.Data=} :: {ex.Len}")
print(f" >> {pp.Data=} :: {pp.Len}")
