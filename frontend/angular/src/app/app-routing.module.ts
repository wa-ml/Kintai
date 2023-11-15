import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserListComponent } from './pages/admin/user-list/user-list.component';
import { UserRegisterComponent } from './pages/admin/user-register/user-register.component';
import { RegisterCompleteComponent } from './pages/admin/register-complete/register-complete.component';
import { LoginComponent } from './pages/login/login.component';
import { AuthGuard } from './auth.guard';

const routes: Routes = [
  { path: "login", component: LoginComponent, title: "研究室入退室管理App" },
  { path: "admin/reporters", component: UserListComponent, canActivate: [AuthGuard], title: "研究室入退室管理App" },
  { path: "admin/register", component: UserRegisterComponent, canActivate: [AuthGuard], title: "研究室入退室管理App" },
  { path: "admin/complete", component: RegisterCompleteComponent, canActivate: [AuthGuard], title: "研究室入退室管理App" }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
