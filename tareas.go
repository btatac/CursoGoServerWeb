package main

import (
	"fmt"
)

type task struct {
	nombre      string
	description string
	complete    bool
}

type taskList struct {
	task []*task
}

func (t *taskList) agregaraALista(tl *task) {
	t.task = append(t.task, tl)

}

func (t *taskList) eliminarDeLista(index int) {
	t.task = append(t.task[:index], t.task[index+1:]...)
}

func (t *task) marcarCompleta() {
	t.complete = true
}

func (t *task) actualizarDescrip(newDesp string) {
	t.description = newDesp
}

func (t *task) actulizarNombre(newName string) {
	t.nombre = newName
}

func (t *taskList) imprimirLista() {
	println("		*** LISTA DE TAREAS ***          ")
	for _, tarea := range t.task {
		fmt.Println("Nombre: ", tarea.nombre)
		fmt.Println("Descripciòn: ", tarea.description)
		fmt.Println("-------------------------------------------------")

	}
}

func (t *taskList) imprimirListaCompletado() {
	println("		*** TAREAS COMPLETADAS ***          ")
	for _, tarea := range t.task {
		if tarea.complete {
			fmt.Println("Nombre: ", tarea.nombre)
			fmt.Println("Descripciòn: ", tarea.description)
			fmt.Println("-------------------------------------------------")
		}

	}
}

func main() {

	t1 := &task{
		nombre:      "completar mi curso phyton",
		description: "completar mi curso de phyton pronto",
	}

	t2 := &task{
		nombre:      "completar mi curso Go ",
		description: "completar mi curso de Go pronto",
	}

	t3 := &task{
		nombre:      "completar mi curso NODE",
		description: "completar mi curso de Node pronto",
	}

	lista := &taskList{
		task: []*task{
			t1, t2,
		},
	}

	lista.agregaraALista(t3)
	//lista.imprimirLista()
	lista.task[0].marcarCompleta()
	//lista.imprimirListaCompletado()

	mapaTareaas := make(map[string]*taskList)
	mapaTareaas["Tatiana"] = lista

	t4 := &task{
		nombre:      "completar mi curso C# ",
		description: "completar mi curso de C# pronto",
	}

	t5 := &task{
		nombre:      "completar mi curso Angular",
		description: "completar mi curso de Angular pronto",
	}

	lista2 := &taskList{
		task: []*task{
			t4, t5,
		},
	}

	mapaTareaas["Gondo Momoso"] = lista2

	fmt.Println("Tareas de Tatiana")
	mapaTareaas["Tatiana"].imprimirLista()

	fmt.Println("Tareas de Gondo")
	mapaTareaas["Gondo Momoso"].imprimirLista()

	/*for i := 0; i < len(lista.task); i++ {
		fmt.Println("Task: ", i, "nombre: ", lista.task[i].nombre)
	}

	/// esta es la mejor forma de usar el ciclo for
	for index, tarea := range lista.task {
		fmt.Println("Task: ", index, "nombre: ", tarea.nombre)
	}

	// Break : rompe el ciclo cuando se usa

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	} */

}
