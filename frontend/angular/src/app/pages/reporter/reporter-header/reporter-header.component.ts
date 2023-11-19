import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthGuard } from '../../../auth.guard';

@Component({
  selector: 'app-reporter-header',
  templateUrl: './reporter-header.component.html',
  styleUrls: ['./reporter-header.component.scss']
})
export class ReporterHeaderComponent {
  constructor(private authGuard: AuthGuard, private router: Router) {}

  logout(): void {
    this.authGuard.logout();
    this.router.navigate(['/login']);
  }
}
