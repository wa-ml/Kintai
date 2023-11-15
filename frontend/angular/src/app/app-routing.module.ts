import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserListComponent } from './pages/admin/user-list/user-list.component';
import { UserRegisterComponent } from './pages/admin/user-register/user-register.component';
import { LoginComponent } from './pages/login/login.component';

const routes: Routes = [
  { path: "login", component: LoginComponent, title: "研究室入退室管理App" },
  { path: "admin/reporters", component: UserListComponent, title: "研究室入退室管理App" },
  { path: "admin/register", component: UserRegisterComponent, title: "研究室入退室管理App" }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
