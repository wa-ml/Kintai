import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReporterHeaderComponent } from './reporter-header.component';

describe('ReporterHeaderComponent', () => {
  let component: ReporterHeaderComponent;
  let fixture: ComponentFixture<ReporterHeaderComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ReporterHeaderComponent]
    });
    fixture = TestBed.createComponent(ReporterHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
