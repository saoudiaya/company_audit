import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-controllist',
  templateUrl: './controllist.component.html',
  styleUrls: ['./controllist.component.css']
})
export class ControllistComponent implements OnInit {

  closeResult: string = '';
  form!:FormGroup;
  link : any;
  list: any[] = [];
  criteria: any[] = [];
  audit: any[] = [];

  
  filteredTeams = [...this.list];
  searchQuery = '';
  
  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.list.filter(team => {
      return (
        team.Name.toLowerCase().includes(filterValue) ||
        team.audit.toLowerCase().includes(filterValue) ||
        team.section.toLowerCase().includes(filterValue) ||
        team.requirement.toLowerCase().includes(filterValue)||
        team.norms.toLowerCase().includes(filterValue)

      );
    });
  }
  constructor(private modalService: NgbModal,private build: FormBuilder,private service: ServiceService,private snackBar: MatSnackBar) { }
  

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      audit: ['', Validators.required],
      description: ['', Validators.required],
      critere: ['', Validators.required],
    });
    this.getControlList();
    this.getAudit();
    this.getCriteria();
  }
open(content:any) {
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
    });
}

open1(content:any) {
this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
  this.closeResult = `Closed with: ${result}`;
}, (reason) => {
  this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
});
}
open2(content:any) {
  this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
    this.closeResult = `Closed with: ${result}`;
  }, (reason) => {
    this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
  });
}
open3(content:any) {
  this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
    this.closeResult = `Closed with: ${result}`;
  }, (reason) => {
    this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
  });
}
private getDismissReason(reason: any): string {
  if (reason === ModalDismissReasons.ESC) {
    return 'by pressing ESC';
  } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
    return 'by clicking on a backdrop';
  } else {
    return  `with: ${reason}`;
  }
}
addControl(){

  const data = {
    nom: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    critere_id: parseInt(this.form.value.critere),
    description: this.form.value.description,
    approbationauditiee: false
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Control List added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addControlList(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getControlList();
      // Reset the form
      this.form.reset();
      // Close the modal
      this.modalService.dismissAll();
    },
    (error) => {
      console.log('Error', error);
    }
  )
}
EditControlList(id : number){

  const data = {
    nom: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    critere_id: parseInt(this.form.value.critere),
    description: this.form.value.description,
    approbationauditiee: false
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Control List edited successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.editControlList(id,data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getControlList();
      // Reset the form
      this.form.reset();
      // Close the modal
      this.modalService.dismissAll();
    },
    (error) => {
      console.log('Error', error);
    }
  )
}
deleteControl(id : number){
  const confirmed = confirm('Are you sure you want to delete this list?');
    if (confirmed) {
      this.snackBar.open('Control List deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      }); 
      this.service.deleteControlList(id).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getControlList();
          this.form.reset();
          // Close the modal
          this.modalService.dismissAll();
        },
        (error) => {
          console.log('Error', error);
        }
      );
    }
}
EditControlList2(){
  this.snackBar.open('Note added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
}
getControlList() {
  this.service.getControlList().subscribe(
    (response: any[]) => {
      this.list = response; 
    },
    error => {
      console.log('Error', error);
    }
  );
}
getCriteria() {
  this.service.getCriteria().subscribe(
    (response: any[]) => {
      this.criteria = response; 
    },
    error => {
      console.log('Error', error);
    }
  );
}
getAudit() {
  this.service.getAudit().subscribe(
    (response: any[]) => {
      this.audit = response; 
    },
    error => {
      console.log('Error', error);
    }
  );
}


}
