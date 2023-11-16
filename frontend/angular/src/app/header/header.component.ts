import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthGuard } from '../auth.guard';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent {
  constructor(private authGuard: AuthGuard, private router: Router) {}

  logout(): void {
    this.authGuard.logout();
    this.router.navigate(['/login']);
  }
}
