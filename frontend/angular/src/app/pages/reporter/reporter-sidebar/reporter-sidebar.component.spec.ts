import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReporterSidebarComponent } from './reporter-sidebar.component';

describe('ReporterSidebarComponent', () => {
  let component: ReporterSidebarComponent;
  let fixture: ComponentFixture<ReporterSidebarComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ReporterSidebarComponent]
    });
    fixture = TestBed.createComponent(ReporterSidebarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
