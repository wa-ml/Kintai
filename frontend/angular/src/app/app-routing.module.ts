import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserListComponent } from './pages/admin/user-list/user-list.component';
import { UserRegisterComponent } from './pages/admin/user-register/user-register.component';
import { RegisterCompleteComponent } from './pages/admin/register-complete/register-complete.component';
import { LoginComponent } from './pages/login/login.component';
import { ReporterRecordComponent } from './pages/reporter/reporter-record/reporter-record.component';
import { ReporterRegisterComponent } from './pages/reporter/reporter-register/reporter-register.component';
import { AuthGuard } from './auth.guard';

const routes: Routes = [
  { path: "", component: LoginComponent, title: "研究室入退室管理App" },
  { path: "login", component: LoginComponent, title: "研究室入退室管理App" },
  { path: "record", component: ReporterRecordComponent, canActivate: [AuthGuard], title: "報告者画面" },
  { path: "register", component: ReporterRegisterComponent, canActivate: [AuthGuard], title: "報告者画面" },
  { path: "admin/reporters", component: UserListComponent, canActivate: [AuthGuard], title: "管理者画面" },
  { path: "admin/register", component: UserRegisterComponent, canActivate: [AuthGuard], title: "管理者画面" },
  { path: "admin/complete", component: RegisterCompleteComponent, canActivate: [AuthGuard], title: "管理者画面" }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
