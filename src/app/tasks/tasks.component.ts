import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-tasks',
  templateUrl: './tasks.component.html',
  styleUrls: ['./tasks.component.css']
})
export class TasksComponent implements OnInit {

  closeResult: string = '';
  form!:FormGroup;
  link : any;
  task: any[] = [];


  filteredTeams = [...this.task];
  searchQuery = '';
  
  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.task.filter(team => {
      return (
        team.Name.toLowerCase().includes(filterValue) ||
        team.user.toLowerCase().includes(filterValue) ||
        team.audit.toLowerCase().includes(filterValue)  ||
        team.deadline.toLowerCase().includes(filterValue)        
        );
    });
  }
  constructor(private modalService: NgbModal,private build: FormBuilder,private service: ServiceService,private snackBar: MatSnackBar) { }
  

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      user: ['', Validators.required],
      deadline: ['', [Validators.required, Validators.email]],
      audit: ['', Validators.required],
      description: ['', Validators.required],
    });
    this.getAudit();
    this.getTeams();
    this.getTask();
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
addTask(){

  const data = {
    nom: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    utilisateur_id: parseInt(this.form.value.user),
    datefin: this.form.value.deadline.toString(),
    description: this.form.value.description
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Task added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addTask(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getTask();
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
EditTask(id: number){
  const data = {
    nom: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    utilisateur_id: parseInt(this.form.value.user),
    datefin: this.form.value.deadline.toString(),
    description: this.form.value.description
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Task edited successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.editTask(id,data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getTask();
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
deleteTask(id: number){
  const confirmed = confirm('Are you sure you want to delete this task?');
    if (confirmed) {
      this.snackBar.open('Task deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      }); 
      this.service.deleteTask(id).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getTask();
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
getTask() {
  this.service.getTask().subscribe(
    (response: any[]) => {
      this.task = response; 
    },
    error => {
      console.log('Error', error);
    }
  );
}
team: any[] = [];

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
}
