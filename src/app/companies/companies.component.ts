import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-companies',
  templateUrl: './companies.component.html',
  styleUrls: ['./companies.component.css']
})
export class CompaniesComponent implements OnInit {
  closeResult: string = '';
  form!: FormGroup;
  company: any[] = [];
  filteredTeams = [...this.company];
  searchQuery = '';

  constructor(
    private modalService: NgbModal,
    private build: FormBuilder,
    private service: ServiceService,
    private snackBar: MatSnackBar
  ) {}

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      address: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      phone: ['', Validators.required]
    });
    this.getCompany();
  }

  open(content: any) {
    this.modalService.open(content, { ariaLabelledBy: 'modal-basic-title' }).result.then(
      (result) => {
        this.closeResult = `Closed with: ${result}`;
      },
      (reason) => {
        this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
      }
    );
  }

  open1(content: any) {
    this.modalService.open(content, { ariaLabelledBy: 'modal-basic-title' }).result.then(
      (result) => {
        this.closeResult = `Closed with: ${result}`;
      },
      (reason) => {
        this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
      }
    );
  }

  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return 'by clicking on a backdrop';
    } else {
      return `with: ${reason}`;
    }
  }

  getCompany() {
    this.service.getEntreprise().subscribe(
      (response: any[]) => {
        this.company = response;
        this.filteredTeams = [...this.company];
      },
      (error) => {
        console.log('Error', error);
      }
    );
  }

  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.company.filter((team) => {
      return (
        team.nom.toLowerCase().includes(filterValue) ||
        team.email.toLowerCase().includes(filterValue) ||
        team.phone.toLowerCase().includes(filterValue) ||
        team.address.toLowerCase().includes(filterValue)
      );
    });
  }

  addCompany() {
    if (this.form.valid) {
      const data = {
        ...this.form.value,
        managed_by: 1
      };
  
      console.log('Data:', data); // Display the data in the console
      this.snackBar.open('Company added successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      });      
      this.service.addEntreprise(data).subscribe(
        (response: any) => {
          // Refresh the company list after successful addition
          this.getCompany();
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
  }
  

  editCompany(id:any) {
    if (this.form.valid) {
      const data = {
        ...this.form.value,
        managed_by: 1
      };
      console.log('Data:', data);   
      const companyId = id 
      this.snackBar.open('Company edited successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      });// Replace with the ID of the company being edited
      this.service.editEntreprise(companyId, data).subscribe(
        (response: any) => {
          // Refresh the company list after successful edit
          this.getCompany();
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
  }

  deleteCompany(companyId: any) {
    const confirmed = confirm('Are you sure you want to delete this company?');
    if (confirmed) {
      this.snackBar.open('Company deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      });
      this.service.deleteEntreprise(companyId).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getCompany();
        },
        (error) => {
          console.log('Error', error);
        }
      );
    }
  }
  
}
