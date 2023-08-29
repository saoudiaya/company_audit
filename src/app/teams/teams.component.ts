import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ServiceService } from '../service.service';
import { MatSnackBar } from '@angular/material/snack-bar';


@Component({
  selector: 'app-teams',
  templateUrl: './teams.component.html',
  styleUrls: ['./teams.component.css']
})
export class TeamsComponent implements OnInit {
  constructor(private modalService: NgbModal, private build: FormBuilder, private service: ServiceService,private snackBar: MatSnackBar) {}
  closeResult: string = '';
  form!: FormGroup;
  link: any;
  team: any[] = [];
  company: any[] = [];
  role: any[] = [];

  filteredTeams = [...this.team];
  searchQuery = '';

  searchTable() {
    const filterValue = this.searchQuery.toLowerCase();

    this.filteredTeams = this.team.filter(t => {
      return (
        t.nom.toLowerCase().includes(filterValue) ||
        t.entreprise.nom.toLowerCase().includes(filterValue) ||
        t.role.nom.toLowerCase().includes(filterValue)
      );
    });
  }

  ngOnInit(): void {
    this.form = this.build.group({
      nom: ['', Validators.required],
      address: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      phone: ['', Validators.required],
      password: ['', Validators.required],
      role_nom: ['', Validators.required],
      entreprise: ['', Validators.required]
    });
    
    
    this.getTeams();
    this.getCompany();
    this.getRole();
  }

  open(content: any) {
    this.modalService.open(content, { ariaLabelledBy: 'modal-basic-title' }).result.then(
      result => {
        this.closeResult = `Closed with: ${result}`;
      },
      reason => {
        this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
      }
    );
  }

  open1(content: any) {
    this.modalService.open(content, { ariaLabelledBy: 'modal-basic-title' }).result.then(
      result => {
        this.closeResult = `Closed with: ${result}`;
      },
      reason => {
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

  getRole() {
    this.service.getRole().subscribe(
      (response: any[]) => {
        this.role = response;
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

  addTeams() {
    if (this.form.valid) {
      const selectedRole = this.role.find(r => r.nom === this.form.value.role_nom);
      const formData = {
        nom: this.form.value.nom,
        address: this.form.value.address,
        email: this.form.value.email,
        phone: this.form.value.phone,
        password: this.form.value.password, // Use plain text password
        role_nom: this.form.value.role_nom,
        role_id: selectedRole ? selectedRole.id : null,
        entreprise_id: parseInt(this.form.value.entreprise, 10)
      };
  
      // TODO: Send formData to the backend
      console.log(formData);
      this.snackBar.open('User added successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      });
            this.service.addUtilisateur(formData).subscribe(
        response => {
          // Handle the successful addition of the user
          console.log('User added successfully', response);
        },
        error => {
          // Handle the error if the addition of the user fails
          console.log('Error adding user', error);
        }
      );
      
    }
  }
  
  EditTeams(id: any) {
    if (this.form.valid) {
      const data = this.form.value;
  
      const teamToUpdate = this.team.find(t => t.id === id);
      const selectedRole = this.role.find(r => r.nom === this.form.value.role_nom);
  
      if (teamToUpdate) {
        teamToUpdate.nom = data.nom;
        teamToUpdate.address = data.address;
        teamToUpdate.email = data.email;
        teamToUpdate.phone = data.phone;
        teamToUpdate.password = data.password; // Use plain text password
        teamToUpdate.role_nom = data.role_nom;
        teamToUpdate.role_id = selectedRole ? selectedRole.id : null; 
        teamToUpdate.entreprise_id = parseInt(data.entreprise, 10);
  
        console.log('Data to be sent:', teamToUpdate);
        this.snackBar.open('User updated successfully', 'Close', { 
          duration: 3000,
          panelClass: 'custom-toast'
        });
        // TODO: Send updated team data to the backend
        this.service.editUtilisateur(id,teamToUpdate).subscribe(
          response => {
            // Handle the successful update of the team
            console.log('Team updated successfully', response);

          },
          error => {
            // Handle the error if the update of the team fails
            console.log('Error updating team', error);
          }
        );
      }
    }
  }
  
  
  
  deleteTeam(id: any) {
    const confirmed = confirm('Are you sure you want to delete this user?');
    if (confirmed) {
      this.snackBar.open('User deleted successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      });
      this.service.deleteUtilisateur(id).subscribe(
        (response: any) => {
          // Refresh the company list after successful deletion
          this.getCompany();  
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
  
  

