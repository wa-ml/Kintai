import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  private loginUrl = 'http://localhost:8080/login';

  loginData = {
    email: "",
    password: ""
  }

  constructor(private http: HttpClient, private router: Router) {}
  
  login(): void {
    const { email, password } = this.loginData;

    this.http.post<any>(this.loginUrl, { "Email": email, "Password": password})
    .subscribe((response) => {
      const token = response.token;
      console.log(token);
      
    });
  }
}
