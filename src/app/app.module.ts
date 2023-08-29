import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SigninComponent } from './signin/signin.component';
import { SignupComponent } from './signup/signup.component';
import { AccueilComponent } from './accueil/accueil.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { TeamsComponent } from './teams/teams.component';
import { CompaniesComponent } from './companies/companies.component';
import { AuditComponent } from './audit/audit.component';
import { ControllistComponent } from './controllist/controllist.component';
import { DocumentsComponent } from './documents/documents.component';
import { MeetingComponent } from './meeting/meeting.component';
import { TasksComponent } from './tasks/tasks.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { PermissionComponent } from './permission/permission.component';
import { RequirementComponent } from './requirement/requirement.component';
import { PerimeterComponent } from './perimeter/perimeter.component';
import { ObjectiveComponent } from './objective/objective.component';
import { CriteriaComponent } from './criteria/criteria.component';
import { NormComponent } from './norm/norm.component';
import { HttpClientModule } from '@angular/common/http';
import { ServiceService } from './service.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatDialogModule } from '@angular/material/dialog';
import { MatSnackBarModule } from '@angular/material/snack-bar';

@NgModule({
  declarations: [
    AppComponent,
    SigninComponent,
    SignupComponent,
    AccueilComponent,
    DashboardComponent,
    TeamsComponent,
    CompaniesComponent,
    AuditComponent,
    ControllistComponent,
    DocumentsComponent,
    MeetingComponent,
    TasksComponent,
    PermissionComponent,
    RequirementComponent,
    PerimeterComponent,
    ObjectiveComponent,
    CriteriaComponent,
    NormComponent,
      ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatDialogModule,
    BrowserAnimationsModule,
    MatSnackBarModule,
  ],
  providers: [ServiceService],
  bootstrap: [AppComponent]
})
export class AppModule { }
