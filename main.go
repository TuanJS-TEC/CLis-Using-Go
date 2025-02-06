package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	todos.add("buy milk")
	todos.add("buy bread")
	todos.toggle(0)
	todos.print()
	storage.Save(todos)
}
