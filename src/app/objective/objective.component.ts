import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-objective',
  templateUrl: './objective.component.html',
  styleUrls: ['./objective.component.css']
})
export class ObjectiveComponent {
  closeResult: string = '';
  form!:FormGroup;
  link : any;
  objectif: any[] = [];

 

  filteredTeams = [...this.objectif];
  searchQuery = '';
  
  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.objectif.filter(team => {
      return (
        team.Name.toLowerCase().includes(filterValue) ||
        team.interview.toLowerCase().includes(filterValue) ||
        team.analysis.toLowerCase().includes(filterValue)  ||
        team.norm.toLowerCase().includes(filterValue) ||
        team.audit.toLowerCase().includes(filterValue) 


        );
    });
  }
  constructor(private modalService: NgbModal,private build: FormBuilder,private service: ServiceService,private snackBar: MatSnackBar) { }
audit: any[] = [];

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
getObjective() {
  this.service.getObjective().subscribe(
    (response: any[]) => {
      this.objectif = response; 
    },
    error => {
      console.log('Error', error);
    }
  );
}

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      description: ['', Validators.required],
      audit: ['', Validators.required],
      });
    this.getAudit();    
    this.getObjective();
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
private getDismissReason(reason: any): string {
  if (reason === ModalDismissReasons.ESC) {
    return 'by pressing ESC';
  } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
    return 'by clicking on a backdrop';
  } else {
    return  `with: ${reason}`;
  }
}
addObjective(){
  const data = {
    nom: this.form.value.nom,
    description: this.form.value.description,
    audit_id:parseInt(this.form.value.audit),
    
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Objective added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addObjective(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getObjective();
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
EditObjective(id : number){
  const data = {
    nom: this.form.value.nom,
    description: this.form.value.description,
    audit_id:parseInt(this.form.value.audit),
    
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Objectif edited successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.editObjective(id,data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getObjective();
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
deleteObjective(id : number){
  const confirmed = confirm('Are you sure you want to delete this objectif?');
    if (confirmed) {
      this.snackBar.open('Objectif deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      }); 
      this.service.deleteObjective(id).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getObjective();
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



}
