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
  // public completedList$$ = new BehaviorSubject<Todo[]>([]);
  // public incompleteList$$ = new BehaviorSubject<Todo[]>([]);
  public todoList$$ = new BehaviorSubject<Todo[]>([]);

  constructor(private todoService: TodoService) { }

  ngOnInit(): void {
    this.getTodoList();
  }

  private async getTodoList() {
    const todoList = await this.todoService.getTodoList().toPromise();
    this.todoList$$.next(todoList);
    // this.completedList$$.next(todoList.filter(function (todo) {
    //   return todo.complete === true;
    // }));
    // this.incompleteList$$.next(todoList.filter(function (todo) {
    //   return todo.complete === false;
    // }));
  }

  // completeTodo(todo: Todo) {
  //   this.todoService.completeTodo(todo).subscribe(() => {
  //     this.getTodoList();
  //   });
  // }

  toggleTodo(todo: Todo) {
    if(todo.complete) {
      this.todoService.updateTodo(todo).subscribe(() => {
        this.getTodoList();
      });
    }
  }

  deleteTodo(todo: Todo) {
    this.todoService.deleteTodo(todo).subscribe(() => {
      this.getTodoList();
    });
  }

}
