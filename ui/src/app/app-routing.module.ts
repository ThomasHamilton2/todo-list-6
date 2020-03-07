import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
// import * as fromRouterConstants from './app-routing.constants';
import * as fromComponents from './components';

const routes: Routes = [
  // {
  //   component: fromComponents.TodoComponent,
  //   path: `todo/:${fromRouterConstants.POKEMON_ID}`,
  // },
  {
    path: 'todo',
    component: fromComponents.TodoComponent,
  },
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'todo',
  },
];

@NgModule({
  exports: [RouterModule],
  imports: [RouterModule.forRoot(routes)],
})

export class AppRoutingModule { }
