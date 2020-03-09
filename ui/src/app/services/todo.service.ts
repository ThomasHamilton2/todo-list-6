import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Todo } from '../models/todo'

@Injectable()
export class TodoService {
  todoUrl: string = '/todo'
  constructor(private http: HttpClient) {}

  getTodoList() {
    return this.http.get<Todo[]>(environment.gateway + this.todoUrl);
  }

  addTodo(todo: Todo) {
    return this.http.post(environment.gateway + this.todoUrl, todo);
  }

  updateTodo(todo: Todo) {
    return this.http.put<number>(environment.gateway + this.todoUrl, todo);
  }

  deleteTodo(todo: Todo) {
    return this.http.delete(environment.gateway + this.todoUrl + '?id=' + todo.id);
  }
}