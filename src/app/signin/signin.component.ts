import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.css']
})
export class SigninComponent implements OnInit {
  currentImageIndex: number = 1;
  currentImage: string = `../../assets/bg/image${this.currentImageIndex}.jpg`;
  currentText: string = 'Create your own courses';

  getImageClass(imageIndex: number): string {
    return `image img-${imageIndex} ${imageIndex === this.currentImageIndex ? 'show' : ''}`;
  }

  selectImage(imageIndex: number): void {
    this.currentImageIndex = imageIndex;
    this.currentImage = `../../assets/bg/image${imageIndex}.jpg`;

    switch (imageIndex) {
      case 1:
        this.currentText = 'Create your own courses';
        break;
      case 2:
        this.currentText = 'Customize as you like';
        break;
      case 3:
        this.currentText = 'Invite students to your class';
        break;
      default:
        this.currentText = '';
    }
  }
  signin(): void {
    // Perform sign-in logic here
    if (this.email === 'root@test.com' && this.password === 'aD@4ykJiDjFlJJP0i8iSe6A7*&') {
      // Successful sign-in
      this.snackBar.open(' Sigin successfully', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      });
      this.router.navigate(['/accueil']);
      // Redirect to the desired page or perform other actions
    } else {
      // Failed sign-in
      this.snackBar.open(' Sigin Failed', 'Close', { 
        duration: 3000,
        panelClass: 'custom-toast'
      });
      // Display an error message or perform other actions
    }
  }
  constructor(private snackBar: MatSnackBar,private router: Router) { }
  email: string | undefined;
  password: string | undefined;
  
  ngOnInit(): void {
  }

}
