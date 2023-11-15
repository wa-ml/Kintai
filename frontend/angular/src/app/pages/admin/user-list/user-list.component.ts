import { Component, OnInit } from '@angular/core';
import { ReporterService } from 'src/app/services/reporter.service';
import { Reporter } from 'src/app/types/reporter';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss']
})
export class UserListComponent implements OnInit {
  reporters: Reporter[] = [];
  currentPage: number = 1;
  itemsPerPage: number = 5;

  constructor(private reporterService: ReporterService) {}

  ngOnInit(): void {
    this.getReporter();
  }

  getReporter(): void {
    this.reporterService.getReporters().subscribe((reporters) => {
      this.reporters = reporters;
    });
  }

  getPageItems(): Reporter[] {
    const startIndex = (this.currentPage - 1) * this.itemsPerPage;
    const endIndex = startIndex + this.itemsPerPage;
    return this.reporters.slice(startIndex, endIndex);
  }

  onPageChange(pageNumber: number): void {
    this.currentPage = pageNumber;
  }

  totalPages(): number {
    return Math.ceil(this.reporters.length / this.itemsPerPage);
  }

  getPageNumbers(): number[] {
    const pageCount = this.totalPages();
    return Array.from({ length: pageCount }, (_, index) => index + 1);
  }
}
