import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-meeting',
  templateUrl: './meeting.component.html',
  styleUrls: ['./meeting.component.css']
})
export class MeetingComponent implements OnInit {

  closeResult: string = '';
  form!:FormGroup;
  link : any;
  
  meet: any[] = [];

  

  filteredTeams = [...this.meet];
  searchQuery = '';
  
  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.meet.filter(team => {
      return (
        team.Name.toLowerCase().includes(filterValue) ||
        team.audit.toLowerCase().includes(filterValue)
      );
    });
  }
  constructor(private modalService: NgbModal,private build: FormBuilder,private service: ServiceService,private snackBar: MatSnackBar) { }
  

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      audit: ['', Validators.required],
      user: ['', Validators.required],
      starttime: ['', Validators.required],
      endtime: ['', Validators.required],
    });
    this.getAudit();
    this.getMeeting();
    this.getTeams();
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
addMeeting(){

  const data = {
    titre: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    utilisateur_id:parseInt(this.form.value.user),
    date_debut: this.form.value.starttime,
    date_fin:this.form.value.endtime,
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Meeting added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addMeeting(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getMeeting();
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
EditMeeting(id: number){
  const data = {
    titre: this.form.value.nom,
    audit_id: parseInt(this.form.value.audit),
    utilisateur_id:parseInt(this.form.value.user),
    date_debut: this.form.value.starttime.toString(),
    date_fin:this.form.value.endtime.toString(),
  };
  

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Meeting edited successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.editMeeting(id,data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getMeeting();
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
EditMeeting2(id: number){
  
  // Show the data before posting
  this.snackBar.open('Note added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });   
}

deleteMeeting(id: number){
  const confirmed = confirm('Are you sure you want to delete this meet?');
    if (confirmed) {
      this.snackBar.open('Meeting deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      }); 
      this.service.deleteMeeting(id).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getMeeting();
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

getMeeting() {
  this.service.getMeeting().subscribe(
    (response: any[]) => {
      this.meet = response; 
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
getStatus(starttime: string, endtime: string): string {
  const now = new Date();
  const start = new Date(starttime);
  const end = new Date(endtime);

  if (now >= start && now <= end) {
    return 'In Progress';
  } else if (now > end) {
    return 'Finished';
  } else if (now < start) {
    if (start.getDate() === now.getDate()) {
      return 'Not Yet';
    } else {
      return 'Unknown';
    }
  } else {
    return 'Unknown';
  }
}



}