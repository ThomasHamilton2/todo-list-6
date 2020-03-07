import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { TodoComponent } from './components/todo/todo.component';

import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
// import { AppRoutingModule } from './app-routing.module';
import { Route, RouterModule } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';

const ROUTES: Route[] = [
  { path: '', component: TodoComponent},
  { path: 'todo', component: TodoComponent},
  // { path: 'home', component: HomeComponent}
]

@NgModule({
  declarations: [
    AppComponent,
    TodoComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatCardModule,
    MatProgressSpinnerModule,
    // AppRoutingModule,
    RouterModule.forRoot(ROUTES),
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
