package main

func main(){
va := 0
for i := 0;i<10000;i++ {
	go func(){
		va++
	
	}()
}

print(va)
select{}
}
