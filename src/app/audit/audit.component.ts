import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { jsPDF } from 'jspdf';
import html2canvas from 'html2canvas';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-audit',
  templateUrl: './audit.component.html',
  styleUrls: ['./audit.component.css']
})
export class AuditComponent implements OnInit {

  closeResult: string = '';
  form!:FormGroup;
  formR!:FormGroup;
  selectedRole: string | undefined;
  link : any;
  audit: any[] = [];
  team: any[] = [];
  company: any[] = [];

  filteredTeams = [...this.audit];
  searchQuery = '';
  
  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.audit.filter(team => {
      return (
        team.Name.toLowerCase().includes(filterValue) ||
        team.company.toLowerCase().includes(filterValue) ||
        team.startday.toLowerCase().includes(filterValue) ||
        team.endday.toLowerCase().includes(filterValue)
      );
    });
  }
  constructor(private modalService: NgbModal,private build: FormBuilder,private service: ServiceService,private snackBar: MatSnackBar) { }
  

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      company: ['', Validators.required],
      user: ['', [Validators.required]],
      startday: ['', Validators.required],
      endday: ['', Validators.required],
      description: ['', Validators.required],
      note: ['', Validators.required],
    });
    this.formR = this.build.group({
      number: ['', Validators.required],
      process: ['', Validators.required],
      numberA: ['', [Validators.required]],
      description: ['', Validators.required],
      level: ['', Validators.required],

    });
    this.getAudit();
    this.getCompany();
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
open4(content:any) {
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
getCompany() {
  this.service.getEntreprise().subscribe(
    (response: any[]) => {
      this.company = response; 
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

isFinished(endDay: string): boolean {
  const today = new Date();
  const endDate = new Date(endDay);

  // Compare the dates by comparing the year, month, and day
  return endDate.getFullYear() <= today.getFullYear() &&
    endDate.getMonth() <= today.getMonth() &&
    endDate.getDate() <= today.getDate();
}
submitReport(){

  const data = {
    nom: "Rapport",
    description: this.formR.value.description,
    nombreofconformite : this.formR.value.number,
    process: this.formR.value.process,
    nombreofarticle :this.formR.value.numberA,
    observation_id: 1,
  };

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Rapport added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addRapport(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getAudit();
      // Reset the form
      this.form.reset();
      // Close the modal
      this.modalService.dismissAll();
    },
    (error) => {
      console.log('Error', error);
    }
  );
}
addAudit() {
  const selectedUser = this.team.find(r => r.nom === this.form.value.user);
  const selectedCompany = this.company.find(r => r.nom === this.form.value.company);

  // Retrieve the user ID based on the selected user
  const utilisateur_principale = selectedUser ? selectedUser.id : null;

  // Retrieve the company ID based on the selected company
  const entreprise_id = selectedCompany ? selectedCompany.id : null;
  console.log(this.form.value.startday);
  const data = {
    nom: this.form.value.nom,
    description: this.form.value.description,
    type: "Financial Audit",
    statut: "non",
    date_debut: this.form.value.startday,
    date_fin: this.form.value.endday,
    efficacement: false,
    efficacement_jours: null,
    observation: "",
    utilisateur_principale,
    entreprise_auditie: 7,
    entreprise_auditrice: 10,
    entreprise_id
  };

  console.log(data);
  // Show the data before posting
  this.snackBar.open('Audit added successfully', 'Close', { 
    duration: 3000,
    panelClass: 'custom-toast'
  });  
  
  this.service.addAudit(data).subscribe(
    (response: any) => {
      // Refresh the audit list after successful addition
      this.getAudit();
      // Reset the form
      this.form.reset();
      // Close the modal
      this.modalService.dismissAll();
    },
    (error) => {
      console.log('Error', error);
    }
  );
}

deleteAudit(id: any){
  const confirmed = confirm('Are you sure you want to delete this audit?');
    if (confirmed) {
      this.snackBar.open('Audit deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      }); 
      this.service.deleteAudit(id).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getAudit();
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

EditAudit(id : any) {
  const selectedUser = this.team.find(r => r.nom === this.form.value.user);
  const selectedCompany = this.company.find(r => r.nom === this.form.value.company);

  // Retrieve the user ID based on the selected user
  const utilisateur_principale = selectedUser ? selectedUser.id : null;

  // Retrieve the company ID based on the selected company
  const entreprise_id = selectedCompany ? selectedCompany.id : null;
  console.log(this.form.value.startday);
  const data = {
    nom: this.form.value.nom,
    description: this.form.value.description,
    type: "Financial Audit",
    statut: "non",
    date_debut: this.form.value.startday,
    date_fin: this.form.value.endday,
    efficacement: false,
    efficacement_jours: null,
    observation: "",
    utilisateur_principale,
    entreprise_auditie: 7,
    entreprise_auditrice: 10,
    entreprise_id
  };
    this.snackBar.open('Audit edited successfully', 'Close', { 
      duration: 3000,
      panelClass: 'custom-toast'
    }); 
    this.service.editAudit(id, data).subscribe(
      response => {
        // Handle successful response
        console.log('Audit edited successfully.');
        // You can perform any additional actions here, such as displaying a success message or navigating to another page
      },
      error => {
        // Handle error response
        console.error('Failed to edit audit:', error);
        // You can display an error message or perform any other error handling logic here
      }
    );
  }
EditAudit2(auditId: number) {
    const data = {
      nom: this.audit[0].nom,
    description: this.audit[0].description,
    type: "Financial Audit",
    statut: "non",
    date_debut: this.audit[0].startday,
    date_fin: this.audit[0].endday,
    efficacement: false,
    efficacement_jours: null,
    utilisateur_principale:this.audit[0].utilisateur_principale,
    entreprise_auditie: 7,
    entreprise_auditrice: 10,
    entreprise_id:this.audit[0].entreprise_id,
    observation: this.form.value.note
    };
    console.log(data);
    this.snackBar.open('Note added successfully', 'Close', { 
      duration: 3000,
      panelClass: 'custom-toast'
    }); 
    this.service.editAudit(auditId, data).subscribe(
      response => {
        // Handle successful response
        console.log('Audit edited successfully.');
        // You can perform any additional actions here, such as displaying a success message or navigating to another page
      },
      error => {
        // Handle error response
        console.error('Failed to edit audit:', error);
        // You can display an error message or perform any other error handling logic here
      }
    );
  }
  
  
generatePDF() {
  const data = this.formR.value;

    // Get the current date
    const currentDate = new Date().toLocaleDateString();

    // Create a new jsPDF instance
    const doc = new jsPDF();

    // Set the title
    doc.setFontSize(16);
    doc.text('Non-Compliance Report', 10, 10);

    // Set the content
    doc.setFontSize(12);
    doc.text(`Date: ${currentDate}`, 10, 20);
    doc.text(`Number of Non-Compliance: ${data.number}`, 10, 30);
    doc.text(`Process: ${data.process}`, 10, 40);
    doc.text(`Number of Article: ${data.numberA}`, 10, 50);
    doc.text(`Description: ${data.description}`, 10, 60);
    doc.text(`Level: ${data.level}`, 10, 70);

    // Save the PDF file
    doc.save('non-compliance-report.pdf');
  }
}



