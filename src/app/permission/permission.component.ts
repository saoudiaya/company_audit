import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-permission',
  templateUrl: './permission.component.html',
  styleUrls: ['./permission.component.css']
})
export class PermissionComponent {
  constructor(private modalService: NgbModal, private build: FormBuilder, private service: ServiceService,private snackBar: MatSnackBar) {}
  role: any[] = [];

  getRole() {
    this.service.getRole().subscribe(
      (response: any[]) => {
        this.role = response;
      },
      error => {
        console.log('Error', error);
      }
    );
  }
  ngOnInit(): void {

    this.getRole();
  }

selectRole(role: any) {
  this.selectedRole = role;
}
addpermission(){
  this.snackBar.open('Permission added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
}
  selectedRole: string | undefined;
  teamsReadPermission: boolean | undefined;
  teamsWritePermission: boolean | undefined;

  permissionReadPermission: boolean | undefined;
  permissionWritePermission: boolean | undefined;

  companiesReadPermission: boolean | undefined;
  companiesWritePermission: boolean | undefined;

  auditReadPermission: boolean | undefined;
  auditWritePermission: boolean | undefined;

  controllistReadPermission: boolean | undefined;
  controllistWritePermission: boolean | undefined;

  documentsReadPermission: boolean | undefined;
  documentsWritePermission: boolean | undefined;

  meetingReadPermission: boolean | undefined;
  meetingWritePermission: boolean | undefined;

  tasksReadPermission: boolean | undefined;
  tasksWritePermission: boolean | undefined;

  requirementReadPermission: boolean | undefined;
  requirementWritePermission: boolean | undefined;

  perimeterReadPermission: boolean | undefined;
  perimeterWritePermission: boolean | undefined;

  objectivesReadPermission: boolean | undefined;
  objectivesWritePermission: boolean | undefined;

  normReadPermission: boolean | undefined;
  normWritePermission: boolean | undefined;

  criteriaReadPermission: boolean | undefined;
  criteriaWritePermission: boolean | undefined;
} 
