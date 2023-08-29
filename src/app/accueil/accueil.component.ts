import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-accueil',
  templateUrl: './accueil.component.html',
  styleUrls: ['./accueil.component.css']
})
export class AccueilComponent implements OnInit {
  isMenuOpen: boolean = true;
  activePage: string = 'dashboard';
  toggleMenu() {
    this.isMenuOpen = !this.isMenuOpen;
  }

  //Dashboard
  isDashboardVisible = true;
  showDashboard() {
    this.isDashboardVisible = true;
    this.isTeamsVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'dashboard';
  }
  //Teams
  isTeamsVisible = false;

  showTeams() {
    this.isTeamsVisible = true;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'teams';
  }

  //Companies
  isCompaniesVisible = false;
  showCompanies() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = true;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'companies';
  }

  //Audit
  isAuditVisible = false;
  showAudit() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = true;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'audit';
  }

  //ControlList
  isControlListVisible = false;
  showControlList() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = true;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'control';
  }

  //Documents
  isDocumentsVisible = false;
  showDocuments() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = true;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'document';
  }

  //Meeting
  isMeetingVisible = false;
  showMeeting() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = true;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'meeting';
  }

  //Tasks
  isTasksVisible = false;
  showTasks() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = true;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'tasks';
  }

  //Requirement
  isRequirementVisible = false;
  showRequirement() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = true;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'requirement';
  }

  //Perimeter
 isPerimeterVisible = false;
  showPerimeter() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = true;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'perimeter';
  }

  //Objective
  isObjectiveVisible = false;
  showObjective() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = true;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'objective';
  }

  //Norm
  isNormVisible = false;
  showNorm() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = true;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = false;
    this.activePage = 'norm';
  }

  //Criteria
  isCriteriaVisible = false;
  showCriteria() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = true;
    this.isPermissionVisible = false;
    this.activePage = 'criteria';
  }

  //Permission
  isPermissionVisible = false;
  showPermission() {
    this.isTeamsVisible = false;
    this.isDashboardVisible = false;
    this.isCompaniesVisible = false;
    this.isAuditVisible = false;
    this.isControlListVisible = false;
    this.isDocumentsVisible = false;
    this.isMeetingVisible = false;
    this.isTasksVisible = false;
    this.isRequirementVisible = false;
    this.isPerimeterVisible = false;
    this.isObjectiveVisible = false;
    this.isNormVisible = false;
    this.isCriteriaVisible = false;
    this.isPermissionVisible = true;
    this.activePage = 'permission';
  }


  constructor() { }

  ngOnInit(): void {
  }

}
