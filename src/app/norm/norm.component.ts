import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-norm',
  templateUrl: './norm.component.html',
  styleUrls: ['./norm.component.css']
})
export class NormComponent {
  closeResult: string = '';
  form!:FormGroup;
  link : any;
  norm: any[] = [];

 

  filteredTeams = [...this.norm];
  searchQuery = '';
  
  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.norm.filter(team => {
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

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      description: ['', Validators.required],
      audit: ['', Validators.required],
      });
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
addNorm(){
  const data = {
    nom: this.form.value.nom,
    description: this.form.value.description,    
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Norm added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addNorm(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getNorm();
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
EditNorm(id : number){
  const data = {
    nom: this.form.value.nom,
    description: this.form.value.description,
    
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Norm edited successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.editNorm(id,data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getNorm();
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
deleteNorm(id : number){
  const confirmed = confirm('Are you sure you want to delete this norm?');
    if (confirmed) {
      this.snackBar.open('Norm deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      }); 
      this.service.deleteNorm(id).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getNorm();
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
