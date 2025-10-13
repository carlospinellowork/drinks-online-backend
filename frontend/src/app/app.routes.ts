import { Routes } from '@angular/router';
import { Drinks } from './pages/drinks/drinks';
import { Login } from './pages/login/login';

export const routes: Routes = [
  { path: '', component:  Drinks},
  { path: 'login', component: Login}
];
