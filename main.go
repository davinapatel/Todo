package main

func main() {

	actions := ToDos{}
	storage := NewStorage[ToDos]("todos.json")
	storage.Load(&actions)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&actions)
	storage.Save(actions)
}
