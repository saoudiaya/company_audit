import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';


@Component({
  selector: 'app-requirement',
  templateUrl: './requirement.component.html',
  styleUrls: ['./requirement.component.css']
})
export class RequirementComponent {
  closeResult: string = '';
  form!:FormGroup;
  link : any;
  requirement: any[] = [];

 

  filteredTeams = [...this.requirement];
  searchQuery = '';
  
  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.requirement.filter(team => {
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
  

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      audit: ['', Validators.required],
      observation : ['', Validators.required],
      description: ['', Validators.required]
    });
    this.getAudit();
    this.getRequirement();
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
addRequirement(){

  const data = {
    nom: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    description: this.form.value.description,
    niveau: this.form.value.observation
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Requirement added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addRequirement(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getRequirement();
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
getRequirement() {
  this.service.getRequirement().subscribe(
    (response: any[]) => {
      this.requirement = response; 
    },
    error => {
      console.log('Error', error);
    }
  );
}
EditRequirement(id:number){
  const data = {
    nom: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    description: this.form.value.description,
    niveau: this.form.value.observation
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Requirement edited successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.editRequirement(id, data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getRequirement();
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
deleteRequirement(id:number){
  const confirmed = confirm('Are you sure you want to delete this requirement?');
  if (confirmed) {
    this.snackBar.open('Requirement deleted successfully', 'Close', { 
      duration: 3000,
      panelClass: 'custom-toast'
    }); 
    this.service.deleteRequirement(id).subscribe(
      (response: any) => {
        // Refresh the company list after successful deletion
        this.getRequirement();
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
