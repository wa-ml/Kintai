import { Component, OnInit } from '@angular/core';
import { ReporterService } from 'src/app/services/reporter.service';
import { Reporter } from 'src/app/types/reporter';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss']
})
export class UserListComponent  implements OnInit {
  reporters: Reporter[] = [];
  constructor(private reporterService: ReporterService) {}

  ngOnInit(): void {
    this.getReporter();
  }

  getReporter(): void {
    this.reporterService.getReporters().subscribe((reporters) => {
      this.reporters = reporters;
    });
  }
}
