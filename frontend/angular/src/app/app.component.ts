import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = '研究室入退室管理App';
  constructor(private router: Router) {}

  isLoginPage(): boolean {
    return (this.router.url === "/login" || this.router.url === "/");
  }

  isAdminPage(): boolean {
    return this.router.url.includes("/admin");
  }
}
