import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-criteria',
  templateUrl: './criteria.component.html',
  styleUrls: ['./criteria.component.css']
})
export class CriteriaComponent {
  closeResult: string = '';
  form!:FormGroup;
  link : any;

  criteria: any[] = [];

  filteredTeams = [...this.criteria];
  searchQuery = '';
  
  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.criteria.filter(team => {
      return (
        team.Name.toLowerCase().includes(filterValue) ||
        team.interview.toLowerCase().includes(filterValue) ||
        team.analysis.toLowerCase().includes(filterValue)  ||
        team.norm.toLowerCase().includes(filterValue) ||
        team.audit.toLowerCase().includes(filterValue) 


        );
    });
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
norm: any[] = [];

getNorm() {
  this.service.getNorm().subscribe(
    (response: any[]) => {
      this.norm = response; 
    },
    error => {
      console.log('Error', error);
    }
  );
}
constructor(private modalService: NgbModal,private build: FormBuilder,private service: ServiceService,private snackBar: MatSnackBar) { }
  

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      interview: ['', Validators.required],
      norm: ['', Validators.required],
      description: ['', Validators.required],
      analysis: ['', Validators.required],
      technicalV: ['', Validators.required]
    });
    this.getCriteria();
    this.getNorm();
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
addCriteria(){
  const data = {
    nom: this.form.value.nom,
    entretien:this.form.value.interview,
    description: this.form.value.description,
    verification_technique: this.form.value.technicalV,
    norme_id:parseInt(this.form.value.norm),
    analyse: this.form.value.analysis,
    
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Criteria added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addCriteria(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getCriteria();
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
EditCriteria(id : number){
  const data = {
    nom: this.form.value.nom,
    entretien:this.form.value.interview,
    description: this.form.value.description,
    verification_technique: this.form.value.technicalV,
    norme_id:parseInt(this.form.value.norm),
    analyse: this.form.value.analysis,
    
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Criteria edited successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.editCriteria(id,data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getCriteria();
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
deleteCriteria(id : number){
  const confirmed = confirm('Are you sure you want to delete this criteria?');
    if (confirmed) {
      this.snackBar.open('Criteria deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      }); 
      this.service.deleteCriteria(id).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getCriteria();
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
