import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ReporterService } from 'src/app/services/reporter.service';

@Component({
  selector: 'app-user-register',
  templateUrl: './user-register.component.html',
  styleUrls: ['./user-register.component.scss']
})
export class UserRegisterComponent {
  postData = {
    name: "",
    email: ""
  }

  constructor(private reporterSrrvice: ReporterService, private router: Router) {}

  postReporter(): void {    
    let { name, email } = this.postData;
    email = email + "@example.com";

    this.reporterSrrvice.postReporter(name, email).subscribe((response) => {
      this.router.navigate(['/admin/complete']);
    },(err) => {
      console.log(err);
    });
  }
}
