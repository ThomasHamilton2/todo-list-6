import { Component, OnInit } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Todo } from '../../models/todo'
import { TodoService } from '../../services/todo.service'

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.css'],
  providers: [TodoService]
})
export class TodoComponent implements OnInit {
  newTodoTitle: string = '';
  public todoList$$ = new BehaviorSubject<Todo[]>([]);

  constructor(private todoService: TodoService) { }

  ngOnInit(): void {
    this.getTodoList();
  }

  private async getTodoList() {
    const todoList = await this.todoService.getTodoList().toPromise();
    this.todoList$$.next(todoList);
  }

  toggleTodo(todo: Todo) {
    //todo thomas - todo.complete is not updated before this method is called
    //copy into another object and send that one with updated Complete property
    var copy = todo;
    copy.complete = !copy.complete;
    this.todoService.updateTodo(copy).subscribe(() => {
    });
  }

  deleteTodo(todo: Todo) {
    this.todoService.deleteTodo(todo).subscribe(() => {
      this.getTodoList();
    });
  }

  addTodo() {
    var newTodo = new Todo;
    newTodo.complete = false;
    newTodo.title = this.newTodoTitle;
    this.todoService.addTodo(newTodo).subscribe(() => {
      this.newTodoTitle = "";
      this.getTodoList();
    });
  }

}
