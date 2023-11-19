import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  private loginUrl = 'http://localhost:8080/login';
  private admincheckUrl = 'http://localhost:8080/auth/check';

  loginData = {
    email: "",
    password: "",
    errMessage: ""
  }

  constructor(private http: HttpClient, private router: Router) {}
  
  login(): void {
    const { email, password } = this.loginData;

    this.http.post<any>(this.loginUrl, { "Email": email, "Password": password})
    .subscribe((response) => {
      const token = response.token;
      localStorage.setItem('token', token);
      this.http.get<any>(this.admincheckUrl, {
        headers: new HttpHeaders().set("Authorization", `Bearer ${localStorage.getItem('token')}`)
      }).subscribe((response) => {
        // 管理者側のログイン
        if (response) {
          this.router.navigate(['/admin/reporters']);
        // 報告者側のログイン
        } else {
          this.router.navigate(['/reporter']);
        }
      },(err) => {
        this.loginData.errMessage = "サーバ側のエラーが発生しました。"
      })
    },
    (err) => {
      this.loginData.errMessage = "Emailまたはパスワードが異なります。"
      console.log('Login failed', err);
    }
    );
  }
}
