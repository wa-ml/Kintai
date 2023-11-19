import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReporterRegisterComponent } from './reporter-register.component';

describe('ReporterRegisterComponent', () => {
  let component: ReporterRegisterComponent;
  let fixture: ComponentFixture<ReporterRegisterComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ReporterRegisterComponent]
    });
    fixture = TestBed.createComponent(ReporterRegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
