import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-documents',
  templateUrl: './documents.component.html',
  styleUrls: ['./documents.component.css']
})
export class DocumentsComponent implements OnInit {

  closeResult: string = '';
  form!:FormGroup;
  link : any;
  document: any[] = [];
  team: any[] = [];

 

  filteredTeams = [...this.document];
  searchQuery = '';
  
  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.document.filter(team => {
      return (
        team.Name.toLowerCase().includes(filterValue) ||
        team.audit.toLowerCase().includes(filterValue) ||
        team.user.toLowerCase().includes(filterValue)
      );
    });
  }
  constructor(private modalService: NgbModal,private build: FormBuilder,private service: ServiceService,private snackBar: MatSnackBar) { }
  

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      audit: ['', Validators.required],
      user: ['', Validators.required],
    });
    this.getDocument();
    this.getTeams();
    this.getAudit();
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
private getDismissReason(reason: any): string {
  if (reason === ModalDismissReasons.ESC) {
    return 'by pressing ESC';
  } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
    return 'by clicking on a backdrop';
  } else {
    return  `with: ${reason}`;
  }
}
addDocument(){
  const currentDate = new Date().toLocaleDateString();

  const data = {
    titre: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    utilisateur_id: parseInt(this.form.value.user),
    date: currentDate
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Document added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addDocuments(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getDocument();
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
EditDocument(id: number){
  const currentDate = new Date().toLocaleDateString();

  const data = {
    titre: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    utilisateur_id: parseInt(this.form.value.user),
    date: currentDate
  };
    this.snackBar.open('Document edited successfully', 'Close', { 
      duration: 3000,
      panelClass: 'custom-toast'
    }); 
    this.service.editDocuments(id, data).subscribe(
      response => {
       // Refresh the company list after successful deletion
       this.getDocument();
       this.form.reset();
       // Close the modal
       this.modalService.dismissAll();
      },
      error => {
        // Handle error response
        console.error('Failed to edit doc:', error);
        // You can display an error message or perform any other error handling logic here
      }
    );

}

deleteDocument(id: number){
  const confirmed = confirm('Are you sure you want to delete this document?');
    if (confirmed) {
      this.snackBar.open('Document deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      }); 
      this.service.deleteDocuments(id).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getDocument();
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
getImagePath(event:any){
  const path = event.target.files[0].name;
  const link =path;
  this.form.get('image')?.setValue(link);
  console.log(link);
};

getDocument() {
  this.service.getDocuments().subscribe(
    (response: any[]) => {
      this.document = response; 
    },
    error => {
      console.log('Error', error);
    }
  );
}
getTeams() {
  this.service.getUtilisateur().subscribe(
    (response: any[]) => {
      this.team = response; 
    },
    error => {
      console.log('Error', error);
    }
  );
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
addNote(){
  this.snackBar.open('Note added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  }); 
  this.modalService.dismissAll()
}
}
