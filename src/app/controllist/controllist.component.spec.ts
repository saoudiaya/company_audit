import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ControllistComponent } from './controllist.component';

describe('ControllistComponent', () => {
  let component: ControllistComponent;
  let fixture: ComponentFixture<ControllistComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ControllistComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ControllistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
