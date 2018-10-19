// Code generated from pkg/grammars/java/Java9.g4 by ANTLR 4.7.1. DO NOT EDIT.

package java // Java9
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJava9Listener is a complete listener for a parse tree produced by Java9Parser.
type BaseJava9Listener struct{}

var _ Java9Listener = &BaseJava9Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJava9Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJava9Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJava9Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJava9Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJava9Listener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJava9Listener) ExitLiteral(ctx *LiteralContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseJava9Listener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseJava9Listener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterNumericType is called when production numericType is entered.
func (s *BaseJava9Listener) EnterNumericType(ctx *NumericTypeContext) {}

// ExitNumericType is called when production numericType is exited.
func (s *BaseJava9Listener) ExitNumericType(ctx *NumericTypeContext) {}

// EnterIntegralType is called when production integralType is entered.
func (s *BaseJava9Listener) EnterIntegralType(ctx *IntegralTypeContext) {}

// ExitIntegralType is called when production integralType is exited.
func (s *BaseJava9Listener) ExitIntegralType(ctx *IntegralTypeContext) {}

// EnterFloatingPointType is called when production floatingPointType is entered.
func (s *BaseJava9Listener) EnterFloatingPointType(ctx *FloatingPointTypeContext) {}

// ExitFloatingPointType is called when production floatingPointType is exited.
func (s *BaseJava9Listener) ExitFloatingPointType(ctx *FloatingPointTypeContext) {}

// EnterReferenceType is called when production referenceType is entered.
func (s *BaseJava9Listener) EnterReferenceType(ctx *ReferenceTypeContext) {}

// ExitReferenceType is called when production referenceType is exited.
func (s *BaseJava9Listener) ExitReferenceType(ctx *ReferenceTypeContext) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterClassType is called when production classType is entered.
func (s *BaseJava9Listener) EnterClassType(ctx *ClassTypeContext) {}

// ExitClassType is called when production classType is exited.
func (s *BaseJava9Listener) ExitClassType(ctx *ClassTypeContext) {}

// EnterClassType_lf_classOrInterfaceType is called when production classType_lf_classOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterClassType_lf_classOrInterfaceType(ctx *ClassType_lf_classOrInterfaceTypeContext) {
}

// ExitClassType_lf_classOrInterfaceType is called when production classType_lf_classOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitClassType_lf_classOrInterfaceType(ctx *ClassType_lf_classOrInterfaceTypeContext) {
}

// EnterClassType_lfno_classOrInterfaceType is called when production classType_lfno_classOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterClassType_lfno_classOrInterfaceType(ctx *ClassType_lfno_classOrInterfaceTypeContext) {
}

// ExitClassType_lfno_classOrInterfaceType is called when production classType_lfno_classOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitClassType_lfno_classOrInterfaceType(ctx *ClassType_lfno_classOrInterfaceTypeContext) {
}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *BaseJava9Listener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *BaseJava9Listener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterInterfaceType_lf_classOrInterfaceType is called when production interfaceType_lf_classOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterInterfaceType_lf_classOrInterfaceType(ctx *InterfaceType_lf_classOrInterfaceTypeContext) {
}

// ExitInterfaceType_lf_classOrInterfaceType is called when production interfaceType_lf_classOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitInterfaceType_lf_classOrInterfaceType(ctx *InterfaceType_lf_classOrInterfaceTypeContext) {
}

// EnterInterfaceType_lfno_classOrInterfaceType is called when production interfaceType_lfno_classOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterInterfaceType_lfno_classOrInterfaceType(ctx *InterfaceType_lfno_classOrInterfaceTypeContext) {
}

// ExitInterfaceType_lfno_classOrInterfaceType is called when production interfaceType_lfno_classOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitInterfaceType_lfno_classOrInterfaceType(ctx *InterfaceType_lfno_classOrInterfaceTypeContext) {
}

// EnterTypeVariable is called when production typeVariable is entered.
func (s *BaseJava9Listener) EnterTypeVariable(ctx *TypeVariableContext) {}

// ExitTypeVariable is called when production typeVariable is exited.
func (s *BaseJava9Listener) ExitTypeVariable(ctx *TypeVariableContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseJava9Listener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseJava9Listener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterDims is called when production dims is entered.
func (s *BaseJava9Listener) EnterDims(ctx *DimsContext) {}

// ExitDims is called when production dims is exited.
func (s *BaseJava9Listener) ExitDims(ctx *DimsContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseJava9Listener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseJava9Listener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterTypeParameterModifier is called when production typeParameterModifier is entered.
func (s *BaseJava9Listener) EnterTypeParameterModifier(ctx *TypeParameterModifierContext) {}

// ExitTypeParameterModifier is called when production typeParameterModifier is exited.
func (s *BaseJava9Listener) ExitTypeParameterModifier(ctx *TypeParameterModifierContext) {}

// EnterTypeBound is called when production typeBound is entered.
func (s *BaseJava9Listener) EnterTypeBound(ctx *TypeBoundContext) {}

// ExitTypeBound is called when production typeBound is exited.
func (s *BaseJava9Listener) ExitTypeBound(ctx *TypeBoundContext) {}

// EnterAdditionalBound is called when production additionalBound is entered.
func (s *BaseJava9Listener) EnterAdditionalBound(ctx *AdditionalBoundContext) {}

// ExitAdditionalBound is called when production additionalBound is exited.
func (s *BaseJava9Listener) ExitAdditionalBound(ctx *AdditionalBoundContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseJava9Listener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseJava9Listener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterTypeArgumentList is called when production typeArgumentList is entered.
func (s *BaseJava9Listener) EnterTypeArgumentList(ctx *TypeArgumentListContext) {}

// ExitTypeArgumentList is called when production typeArgumentList is exited.
func (s *BaseJava9Listener) ExitTypeArgumentList(ctx *TypeArgumentListContext) {}

// EnterTypeArgument is called when production typeArgument is entered.
func (s *BaseJava9Listener) EnterTypeArgument(ctx *TypeArgumentContext) {}

// ExitTypeArgument is called when production typeArgument is exited.
func (s *BaseJava9Listener) ExitTypeArgument(ctx *TypeArgumentContext) {}

// EnterWildcard is called when production wildcard is entered.
func (s *BaseJava9Listener) EnterWildcard(ctx *WildcardContext) {}

// ExitWildcard is called when production wildcard is exited.
func (s *BaseJava9Listener) ExitWildcard(ctx *WildcardContext) {}

// EnterWildcardBounds is called when production wildcardBounds is entered.
func (s *BaseJava9Listener) EnterWildcardBounds(ctx *WildcardBoundsContext) {}

// ExitWildcardBounds is called when production wildcardBounds is exited.
func (s *BaseJava9Listener) ExitWildcardBounds(ctx *WildcardBoundsContext) {}

// EnterModuleName is called when production moduleName is entered.
func (s *BaseJava9Listener) EnterModuleName(ctx *ModuleNameContext) {}

// ExitModuleName is called when production moduleName is exited.
func (s *BaseJava9Listener) ExitModuleName(ctx *ModuleNameContext) {}

// EnterPackageName is called when production packageName is entered.
func (s *BaseJava9Listener) EnterPackageName(ctx *PackageNameContext) {}

// ExitPackageName is called when production packageName is exited.
func (s *BaseJava9Listener) ExitPackageName(ctx *PackageNameContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseJava9Listener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseJava9Listener) ExitTypeName(ctx *TypeNameContext) {}

// EnterPackageOrTypeName is called when production packageOrTypeName is entered.
func (s *BaseJava9Listener) EnterPackageOrTypeName(ctx *PackageOrTypeNameContext) {}

// ExitPackageOrTypeName is called when production packageOrTypeName is exited.
func (s *BaseJava9Listener) ExitPackageOrTypeName(ctx *PackageOrTypeNameContext) {}

// EnterExpressionName is called when production expressionName is entered.
func (s *BaseJava9Listener) EnterExpressionName(ctx *ExpressionNameContext) {}

// ExitExpressionName is called when production expressionName is exited.
func (s *BaseJava9Listener) ExitExpressionName(ctx *ExpressionNameContext) {}

// EnterMethodName is called when production methodName is entered.
func (s *BaseJava9Listener) EnterMethodName(ctx *MethodNameContext) {}

// ExitMethodName is called when production methodName is exited.
func (s *BaseJava9Listener) ExitMethodName(ctx *MethodNameContext) {}

// EnterAmbiguousName is called when production ambiguousName is entered.
func (s *BaseJava9Listener) EnterAmbiguousName(ctx *AmbiguousNameContext) {}

// ExitAmbiguousName is called when production ambiguousName is exited.
func (s *BaseJava9Listener) ExitAmbiguousName(ctx *AmbiguousNameContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseJava9Listener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseJava9Listener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterOrdinaryCompilation is called when production ordinaryCompilation is entered.
func (s *BaseJava9Listener) EnterOrdinaryCompilation(ctx *OrdinaryCompilationContext) {}

// ExitOrdinaryCompilation is called when production ordinaryCompilation is exited.
func (s *BaseJava9Listener) ExitOrdinaryCompilation(ctx *OrdinaryCompilationContext) {}

// EnterModularCompilation is called when production modularCompilation is entered.
func (s *BaseJava9Listener) EnterModularCompilation(ctx *ModularCompilationContext) {}

// ExitModularCompilation is called when production modularCompilation is exited.
func (s *BaseJava9Listener) ExitModularCompilation(ctx *ModularCompilationContext) {}

// EnterPackageDeclaration is called when production packageDeclaration is entered.
func (s *BaseJava9Listener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {}

// ExitPackageDeclaration is called when production packageDeclaration is exited.
func (s *BaseJava9Listener) ExitPackageDeclaration(ctx *PackageDeclarationContext) {}

// EnterPackageModifier is called when production packageModifier is entered.
func (s *BaseJava9Listener) EnterPackageModifier(ctx *PackageModifierContext) {}

// ExitPackageModifier is called when production packageModifier is exited.
func (s *BaseJava9Listener) ExitPackageModifier(ctx *PackageModifierContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseJava9Listener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseJava9Listener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterSingleTypeImportDeclaration is called when production singleTypeImportDeclaration is entered.
func (s *BaseJava9Listener) EnterSingleTypeImportDeclaration(ctx *SingleTypeImportDeclarationContext) {
}

// ExitSingleTypeImportDeclaration is called when production singleTypeImportDeclaration is exited.
func (s *BaseJava9Listener) ExitSingleTypeImportDeclaration(ctx *SingleTypeImportDeclarationContext) {}

// EnterTypeImportOnDemandDeclaration is called when production typeImportOnDemandDeclaration is entered.
func (s *BaseJava9Listener) EnterTypeImportOnDemandDeclaration(ctx *TypeImportOnDemandDeclarationContext) {
}

// ExitTypeImportOnDemandDeclaration is called when production typeImportOnDemandDeclaration is exited.
func (s *BaseJava9Listener) ExitTypeImportOnDemandDeclaration(ctx *TypeImportOnDemandDeclarationContext) {
}

// EnterSingleStaticImportDeclaration is called when production singleStaticImportDeclaration is entered.
func (s *BaseJava9Listener) EnterSingleStaticImportDeclaration(ctx *SingleStaticImportDeclarationContext) {
}

// ExitSingleStaticImportDeclaration is called when production singleStaticImportDeclaration is exited.
func (s *BaseJava9Listener) ExitSingleStaticImportDeclaration(ctx *SingleStaticImportDeclarationContext) {
}

// EnterStaticImportOnDemandDeclaration is called when production staticImportOnDemandDeclaration is entered.
func (s *BaseJava9Listener) EnterStaticImportOnDemandDeclaration(ctx *StaticImportOnDemandDeclarationContext) {
}

// ExitStaticImportOnDemandDeclaration is called when production staticImportOnDemandDeclaration is exited.
func (s *BaseJava9Listener) ExitStaticImportOnDemandDeclaration(ctx *StaticImportOnDemandDeclarationContext) {
}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseJava9Listener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseJava9Listener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterModuleDeclaration is called when production moduleDeclaration is entered.
func (s *BaseJava9Listener) EnterModuleDeclaration(ctx *ModuleDeclarationContext) {}

// ExitModuleDeclaration is called when production moduleDeclaration is exited.
func (s *BaseJava9Listener) ExitModuleDeclaration(ctx *ModuleDeclarationContext) {}

// EnterModuleDirective is called when production moduleDirective is entered.
func (s *BaseJava9Listener) EnterModuleDirective(ctx *ModuleDirectiveContext) {}

// ExitModuleDirective is called when production moduleDirective is exited.
func (s *BaseJava9Listener) ExitModuleDirective(ctx *ModuleDirectiveContext) {}

// EnterRequiresModifier is called when production requiresModifier is entered.
func (s *BaseJava9Listener) EnterRequiresModifier(ctx *RequiresModifierContext) {}

// ExitRequiresModifier is called when production requiresModifier is exited.
func (s *BaseJava9Listener) ExitRequiresModifier(ctx *RequiresModifierContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseJava9Listener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseJava9Listener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterNormalClassDeclaration is called when production normalClassDeclaration is entered.
func (s *BaseJava9Listener) EnterNormalClassDeclaration(ctx *NormalClassDeclarationContext) {}

// ExitNormalClassDeclaration is called when production normalClassDeclaration is exited.
func (s *BaseJava9Listener) ExitNormalClassDeclaration(ctx *NormalClassDeclarationContext) {}

// EnterClassModifier is called when production classModifier is entered.
func (s *BaseJava9Listener) EnterClassModifier(ctx *ClassModifierContext) {}

// ExitClassModifier is called when production classModifier is exited.
func (s *BaseJava9Listener) ExitClassModifier(ctx *ClassModifierContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseJava9Listener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseJava9Listener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameterList is called when production typeParameterList is entered.
func (s *BaseJava9Listener) EnterTypeParameterList(ctx *TypeParameterListContext) {}

// ExitTypeParameterList is called when production typeParameterList is exited.
func (s *BaseJava9Listener) ExitTypeParameterList(ctx *TypeParameterListContext) {}

// EnterSuperclass is called when production superclass is entered.
func (s *BaseJava9Listener) EnterSuperclass(ctx *SuperclassContext) {}

// ExitSuperclass is called when production superclass is exited.
func (s *BaseJava9Listener) ExitSuperclass(ctx *SuperclassContext) {}

// EnterSuperinterfaces is called when production superinterfaces is entered.
func (s *BaseJava9Listener) EnterSuperinterfaces(ctx *SuperinterfacesContext) {}

// ExitSuperinterfaces is called when production superinterfaces is exited.
func (s *BaseJava9Listener) ExitSuperinterfaces(ctx *SuperinterfacesContext) {}

// EnterInterfaceTypeList is called when production interfaceTypeList is entered.
func (s *BaseJava9Listener) EnterInterfaceTypeList(ctx *InterfaceTypeListContext) {}

// ExitInterfaceTypeList is called when production interfaceTypeList is exited.
func (s *BaseJava9Listener) ExitInterfaceTypeList(ctx *InterfaceTypeListContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseJava9Listener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseJava9Listener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseJava9Listener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseJava9Listener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterClassMemberDeclaration is called when production classMemberDeclaration is entered.
func (s *BaseJava9Listener) EnterClassMemberDeclaration(ctx *ClassMemberDeclarationContext) {}

// ExitClassMemberDeclaration is called when production classMemberDeclaration is exited.
func (s *BaseJava9Listener) ExitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseJava9Listener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseJava9Listener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterFieldModifier is called when production fieldModifier is entered.
func (s *BaseJava9Listener) EnterFieldModifier(ctx *FieldModifierContext) {}

// ExitFieldModifier is called when production fieldModifier is exited.
func (s *BaseJava9Listener) ExitFieldModifier(ctx *FieldModifierContext) {}

// EnterVariableDeclaratorList is called when production variableDeclaratorList is entered.
func (s *BaseJava9Listener) EnterVariableDeclaratorList(ctx *VariableDeclaratorListContext) {}

// ExitVariableDeclaratorList is called when production variableDeclaratorList is exited.
func (s *BaseJava9Listener) ExitVariableDeclaratorList(ctx *VariableDeclaratorListContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseJava9Listener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseJava9Listener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BaseJava9Listener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BaseJava9Listener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseJava9Listener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseJava9Listener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterUnannType is called when production unannType is entered.
func (s *BaseJava9Listener) EnterUnannType(ctx *UnannTypeContext) {}

// ExitUnannType is called when production unannType is exited.
func (s *BaseJava9Listener) ExitUnannType(ctx *UnannTypeContext) {}

// EnterUnannPrimitiveType is called when production unannPrimitiveType is entered.
func (s *BaseJava9Listener) EnterUnannPrimitiveType(ctx *UnannPrimitiveTypeContext) {}

// ExitUnannPrimitiveType is called when production unannPrimitiveType is exited.
func (s *BaseJava9Listener) ExitUnannPrimitiveType(ctx *UnannPrimitiveTypeContext) {}

// EnterUnannReferenceType is called when production unannReferenceType is entered.
func (s *BaseJava9Listener) EnterUnannReferenceType(ctx *UnannReferenceTypeContext) {}

// ExitUnannReferenceType is called when production unannReferenceType is exited.
func (s *BaseJava9Listener) ExitUnannReferenceType(ctx *UnannReferenceTypeContext) {}

// EnterUnannClassOrInterfaceType is called when production unannClassOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterUnannClassOrInterfaceType(ctx *UnannClassOrInterfaceTypeContext) {}

// ExitUnannClassOrInterfaceType is called when production unannClassOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitUnannClassOrInterfaceType(ctx *UnannClassOrInterfaceTypeContext) {}

// EnterUnannClassType is called when production unannClassType is entered.
func (s *BaseJava9Listener) EnterUnannClassType(ctx *UnannClassTypeContext) {}

// ExitUnannClassType is called when production unannClassType is exited.
func (s *BaseJava9Listener) ExitUnannClassType(ctx *UnannClassTypeContext) {}

// EnterUnannClassType_lf_unannClassOrInterfaceType is called when production unannClassType_lf_unannClassOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterUnannClassType_lf_unannClassOrInterfaceType(ctx *UnannClassType_lf_unannClassOrInterfaceTypeContext) {
}

// ExitUnannClassType_lf_unannClassOrInterfaceType is called when production unannClassType_lf_unannClassOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitUnannClassType_lf_unannClassOrInterfaceType(ctx *UnannClassType_lf_unannClassOrInterfaceTypeContext) {
}

// EnterUnannClassType_lfno_unannClassOrInterfaceType is called when production unannClassType_lfno_unannClassOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterUnannClassType_lfno_unannClassOrInterfaceType(ctx *UnannClassType_lfno_unannClassOrInterfaceTypeContext) {
}

// ExitUnannClassType_lfno_unannClassOrInterfaceType is called when production unannClassType_lfno_unannClassOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitUnannClassType_lfno_unannClassOrInterfaceType(ctx *UnannClassType_lfno_unannClassOrInterfaceTypeContext) {
}

// EnterUnannInterfaceType is called when production unannInterfaceType is entered.
func (s *BaseJava9Listener) EnterUnannInterfaceType(ctx *UnannInterfaceTypeContext) {}

// ExitUnannInterfaceType is called when production unannInterfaceType is exited.
func (s *BaseJava9Listener) ExitUnannInterfaceType(ctx *UnannInterfaceTypeContext) {}

// EnterUnannInterfaceType_lf_unannClassOrInterfaceType is called when production unannInterfaceType_lf_unannClassOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterUnannInterfaceType_lf_unannClassOrInterfaceType(ctx *UnannInterfaceType_lf_unannClassOrInterfaceTypeContext) {
}

// ExitUnannInterfaceType_lf_unannClassOrInterfaceType is called when production unannInterfaceType_lf_unannClassOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitUnannInterfaceType_lf_unannClassOrInterfaceType(ctx *UnannInterfaceType_lf_unannClassOrInterfaceTypeContext) {
}

// EnterUnannInterfaceType_lfno_unannClassOrInterfaceType is called when production unannInterfaceType_lfno_unannClassOrInterfaceType is entered.
func (s *BaseJava9Listener) EnterUnannInterfaceType_lfno_unannClassOrInterfaceType(ctx *UnannInterfaceType_lfno_unannClassOrInterfaceTypeContext) {
}

// ExitUnannInterfaceType_lfno_unannClassOrInterfaceType is called when production unannInterfaceType_lfno_unannClassOrInterfaceType is exited.
func (s *BaseJava9Listener) ExitUnannInterfaceType_lfno_unannClassOrInterfaceType(ctx *UnannInterfaceType_lfno_unannClassOrInterfaceTypeContext) {
}

// EnterUnannTypeVariable is called when production unannTypeVariable is entered.
func (s *BaseJava9Listener) EnterUnannTypeVariable(ctx *UnannTypeVariableContext) {}

// ExitUnannTypeVariable is called when production unannTypeVariable is exited.
func (s *BaseJava9Listener) ExitUnannTypeVariable(ctx *UnannTypeVariableContext) {}

// EnterUnannArrayType is called when production unannArrayType is entered.
func (s *BaseJava9Listener) EnterUnannArrayType(ctx *UnannArrayTypeContext) {}

// ExitUnannArrayType is called when production unannArrayType is exited.
func (s *BaseJava9Listener) ExitUnannArrayType(ctx *UnannArrayTypeContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseJava9Listener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BaseJava9Listener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterMethodModifier is called when production methodModifier is entered.
func (s *BaseJava9Listener) EnterMethodModifier(ctx *MethodModifierContext) {}

// ExitMethodModifier is called when production methodModifier is exited.
func (s *BaseJava9Listener) ExitMethodModifier(ctx *MethodModifierContext) {}

// EnterMethodHeader is called when production methodHeader is entered.
func (s *BaseJava9Listener) EnterMethodHeader(ctx *MethodHeaderContext) {}

// ExitMethodHeader is called when production methodHeader is exited.
func (s *BaseJava9Listener) ExitMethodHeader(ctx *MethodHeaderContext) {}

// EnterResult is called when production result is entered.
func (s *BaseJava9Listener) EnterResult(ctx *ResultContext) {}

// ExitResult is called when production result is exited.
func (s *BaseJava9Listener) ExitResult(ctx *ResultContext) {}

// EnterMethodDeclarator is called when production methodDeclarator is entered.
func (s *BaseJava9Listener) EnterMethodDeclarator(ctx *MethodDeclaratorContext) {}

// ExitMethodDeclarator is called when production methodDeclarator is exited.
func (s *BaseJava9Listener) ExitMethodDeclarator(ctx *MethodDeclaratorContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseJava9Listener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseJava9Listener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseJava9Listener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseJava9Listener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseJava9Listener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseJava9Listener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterVariableModifier is called when production variableModifier is entered.
func (s *BaseJava9Listener) EnterVariableModifier(ctx *VariableModifierContext) {}

// ExitVariableModifier is called when production variableModifier is exited.
func (s *BaseJava9Listener) ExitVariableModifier(ctx *VariableModifierContext) {}

// EnterLastFormalParameter is called when production lastFormalParameter is entered.
func (s *BaseJava9Listener) EnterLastFormalParameter(ctx *LastFormalParameterContext) {}

// ExitLastFormalParameter is called when production lastFormalParameter is exited.
func (s *BaseJava9Listener) ExitLastFormalParameter(ctx *LastFormalParameterContext) {}

// EnterReceiverParameter is called when production receiverParameter is entered.
func (s *BaseJava9Listener) EnterReceiverParameter(ctx *ReceiverParameterContext) {}

// ExitReceiverParameter is called when production receiverParameter is exited.
func (s *BaseJava9Listener) ExitReceiverParameter(ctx *ReceiverParameterContext) {}

// EnterThrows_ is called when production throws_ is entered.
func (s *BaseJava9Listener) EnterThrows_(ctx *Throws_Context) {}

// ExitThrows_ is called when production throws_ is exited.
func (s *BaseJava9Listener) ExitThrows_(ctx *Throws_Context) {}

// EnterExceptionTypeList is called when production exceptionTypeList is entered.
func (s *BaseJava9Listener) EnterExceptionTypeList(ctx *ExceptionTypeListContext) {}

// ExitExceptionTypeList is called when production exceptionTypeList is exited.
func (s *BaseJava9Listener) ExitExceptionTypeList(ctx *ExceptionTypeListContext) {}

// EnterExceptionType is called when production exceptionType is entered.
func (s *BaseJava9Listener) EnterExceptionType(ctx *ExceptionTypeContext) {}

// ExitExceptionType is called when production exceptionType is exited.
func (s *BaseJava9Listener) ExitExceptionType(ctx *ExceptionTypeContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseJava9Listener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseJava9Listener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterInstanceInitializer is called when production instanceInitializer is entered.
func (s *BaseJava9Listener) EnterInstanceInitializer(ctx *InstanceInitializerContext) {}

// ExitInstanceInitializer is called when production instanceInitializer is exited.
func (s *BaseJava9Listener) ExitInstanceInitializer(ctx *InstanceInitializerContext) {}

// EnterStaticInitializer is called when production staticInitializer is entered.
func (s *BaseJava9Listener) EnterStaticInitializer(ctx *StaticInitializerContext) {}

// ExitStaticInitializer is called when production staticInitializer is exited.
func (s *BaseJava9Listener) ExitStaticInitializer(ctx *StaticInitializerContext) {}

// EnterConstructorDeclaration is called when production constructorDeclaration is entered.
func (s *BaseJava9Listener) EnterConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// ExitConstructorDeclaration is called when production constructorDeclaration is exited.
func (s *BaseJava9Listener) ExitConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// EnterConstructorModifier is called when production constructorModifier is entered.
func (s *BaseJava9Listener) EnterConstructorModifier(ctx *ConstructorModifierContext) {}

// ExitConstructorModifier is called when production constructorModifier is exited.
func (s *BaseJava9Listener) ExitConstructorModifier(ctx *ConstructorModifierContext) {}

// EnterConstructorDeclarator is called when production constructorDeclarator is entered.
func (s *BaseJava9Listener) EnterConstructorDeclarator(ctx *ConstructorDeclaratorContext) {}

// ExitConstructorDeclarator is called when production constructorDeclarator is exited.
func (s *BaseJava9Listener) ExitConstructorDeclarator(ctx *ConstructorDeclaratorContext) {}

// EnterSimpleTypeName is called when production simpleTypeName is entered.
func (s *BaseJava9Listener) EnterSimpleTypeName(ctx *SimpleTypeNameContext) {}

// ExitSimpleTypeName is called when production simpleTypeName is exited.
func (s *BaseJava9Listener) ExitSimpleTypeName(ctx *SimpleTypeNameContext) {}

// EnterConstructorBody is called when production constructorBody is entered.
func (s *BaseJava9Listener) EnterConstructorBody(ctx *ConstructorBodyContext) {}

// ExitConstructorBody is called when production constructorBody is exited.
func (s *BaseJava9Listener) ExitConstructorBody(ctx *ConstructorBodyContext) {}

// EnterExplicitConstructorInvocation is called when production explicitConstructorInvocation is entered.
func (s *BaseJava9Listener) EnterExplicitConstructorInvocation(ctx *ExplicitConstructorInvocationContext) {
}

// ExitExplicitConstructorInvocation is called when production explicitConstructorInvocation is exited.
func (s *BaseJava9Listener) ExitExplicitConstructorInvocation(ctx *ExplicitConstructorInvocationContext) {
}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseJava9Listener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseJava9Listener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseJava9Listener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseJava9Listener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumConstantList is called when production enumConstantList is entered.
func (s *BaseJava9Listener) EnterEnumConstantList(ctx *EnumConstantListContext) {}

// ExitEnumConstantList is called when production enumConstantList is exited.
func (s *BaseJava9Listener) ExitEnumConstantList(ctx *EnumConstantListContext) {}

// EnterEnumConstant is called when production enumConstant is entered.
func (s *BaseJava9Listener) EnterEnumConstant(ctx *EnumConstantContext) {}

// ExitEnumConstant is called when production enumConstant is exited.
func (s *BaseJava9Listener) ExitEnumConstant(ctx *EnumConstantContext) {}

// EnterEnumConstantModifier is called when production enumConstantModifier is entered.
func (s *BaseJava9Listener) EnterEnumConstantModifier(ctx *EnumConstantModifierContext) {}

// ExitEnumConstantModifier is called when production enumConstantModifier is exited.
func (s *BaseJava9Listener) ExitEnumConstantModifier(ctx *EnumConstantModifierContext) {}

// EnterEnumBodyDeclarations is called when production enumBodyDeclarations is entered.
func (s *BaseJava9Listener) EnterEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// ExitEnumBodyDeclarations is called when production enumBodyDeclarations is exited.
func (s *BaseJava9Listener) ExitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BaseJava9Listener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BaseJava9Listener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterNormalInterfaceDeclaration is called when production normalInterfaceDeclaration is entered.
func (s *BaseJava9Listener) EnterNormalInterfaceDeclaration(ctx *NormalInterfaceDeclarationContext) {}

// ExitNormalInterfaceDeclaration is called when production normalInterfaceDeclaration is exited.
func (s *BaseJava9Listener) ExitNormalInterfaceDeclaration(ctx *NormalInterfaceDeclarationContext) {}

// EnterInterfaceModifier is called when production interfaceModifier is entered.
func (s *BaseJava9Listener) EnterInterfaceModifier(ctx *InterfaceModifierContext) {}

// ExitInterfaceModifier is called when production interfaceModifier is exited.
func (s *BaseJava9Listener) ExitInterfaceModifier(ctx *InterfaceModifierContext) {}

// EnterExtendsInterfaces is called when production extendsInterfaces is entered.
func (s *BaseJava9Listener) EnterExtendsInterfaces(ctx *ExtendsInterfacesContext) {}

// ExitExtendsInterfaces is called when production extendsInterfaces is exited.
func (s *BaseJava9Listener) ExitExtendsInterfaces(ctx *ExtendsInterfacesContext) {}

// EnterInterfaceBody is called when production interfaceBody is entered.
func (s *BaseJava9Listener) EnterInterfaceBody(ctx *InterfaceBodyContext) {}

// ExitInterfaceBody is called when production interfaceBody is exited.
func (s *BaseJava9Listener) ExitInterfaceBody(ctx *InterfaceBodyContext) {}

// EnterInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is entered.
func (s *BaseJava9Listener) EnterInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {}

// ExitInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is exited.
func (s *BaseJava9Listener) ExitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {}

// EnterConstantDeclaration is called when production constantDeclaration is entered.
func (s *BaseJava9Listener) EnterConstantDeclaration(ctx *ConstantDeclarationContext) {}

// ExitConstantDeclaration is called when production constantDeclaration is exited.
func (s *BaseJava9Listener) ExitConstantDeclaration(ctx *ConstantDeclarationContext) {}

// EnterConstantModifier is called when production constantModifier is entered.
func (s *BaseJava9Listener) EnterConstantModifier(ctx *ConstantModifierContext) {}

// ExitConstantModifier is called when production constantModifier is exited.
func (s *BaseJava9Listener) ExitConstantModifier(ctx *ConstantModifierContext) {}

// EnterInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is entered.
func (s *BaseJava9Listener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {}

// ExitInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is exited.
func (s *BaseJava9Listener) ExitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {}

// EnterInterfaceMethodModifier is called when production interfaceMethodModifier is entered.
func (s *BaseJava9Listener) EnterInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) {}

// ExitInterfaceMethodModifier is called when production interfaceMethodModifier is exited.
func (s *BaseJava9Listener) ExitInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) {}

// EnterAnnotationTypeDeclaration is called when production annotationTypeDeclaration is entered.
func (s *BaseJava9Listener) EnterAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {}

// ExitAnnotationTypeDeclaration is called when production annotationTypeDeclaration is exited.
func (s *BaseJava9Listener) ExitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {}

// EnterAnnotationTypeBody is called when production annotationTypeBody is entered.
func (s *BaseJava9Listener) EnterAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// ExitAnnotationTypeBody is called when production annotationTypeBody is exited.
func (s *BaseJava9Listener) ExitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// EnterAnnotationTypeMemberDeclaration is called when production annotationTypeMemberDeclaration is entered.
func (s *BaseJava9Listener) EnterAnnotationTypeMemberDeclaration(ctx *AnnotationTypeMemberDeclarationContext) {
}

// ExitAnnotationTypeMemberDeclaration is called when production annotationTypeMemberDeclaration is exited.
func (s *BaseJava9Listener) ExitAnnotationTypeMemberDeclaration(ctx *AnnotationTypeMemberDeclarationContext) {
}

// EnterAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is entered.
func (s *BaseJava9Listener) EnterAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// ExitAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is exited.
func (s *BaseJava9Listener) ExitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// EnterAnnotationTypeElementModifier is called when production annotationTypeElementModifier is entered.
func (s *BaseJava9Listener) EnterAnnotationTypeElementModifier(ctx *AnnotationTypeElementModifierContext) {
}

// ExitAnnotationTypeElementModifier is called when production annotationTypeElementModifier is exited.
func (s *BaseJava9Listener) ExitAnnotationTypeElementModifier(ctx *AnnotationTypeElementModifierContext) {
}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseJava9Listener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseJava9Listener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseJava9Listener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseJava9Listener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterNormalAnnotation is called when production normalAnnotation is entered.
func (s *BaseJava9Listener) EnterNormalAnnotation(ctx *NormalAnnotationContext) {}

// ExitNormalAnnotation is called when production normalAnnotation is exited.
func (s *BaseJava9Listener) ExitNormalAnnotation(ctx *NormalAnnotationContext) {}

// EnterElementValuePairList is called when production elementValuePairList is entered.
func (s *BaseJava9Listener) EnterElementValuePairList(ctx *ElementValuePairListContext) {}

// ExitElementValuePairList is called when production elementValuePairList is exited.
func (s *BaseJava9Listener) ExitElementValuePairList(ctx *ElementValuePairListContext) {}

// EnterElementValuePair is called when production elementValuePair is entered.
func (s *BaseJava9Listener) EnterElementValuePair(ctx *ElementValuePairContext) {}

// ExitElementValuePair is called when production elementValuePair is exited.
func (s *BaseJava9Listener) ExitElementValuePair(ctx *ElementValuePairContext) {}

// EnterElementValue is called when production elementValue is entered.
func (s *BaseJava9Listener) EnterElementValue(ctx *ElementValueContext) {}

// ExitElementValue is called when production elementValue is exited.
func (s *BaseJava9Listener) ExitElementValue(ctx *ElementValueContext) {}

// EnterElementValueArrayInitializer is called when production elementValueArrayInitializer is entered.
func (s *BaseJava9Listener) EnterElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// ExitElementValueArrayInitializer is called when production elementValueArrayInitializer is exited.
func (s *BaseJava9Listener) ExitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// EnterElementValueList is called when production elementValueList is entered.
func (s *BaseJava9Listener) EnterElementValueList(ctx *ElementValueListContext) {}

// ExitElementValueList is called when production elementValueList is exited.
func (s *BaseJava9Listener) ExitElementValueList(ctx *ElementValueListContext) {}

// EnterMarkerAnnotation is called when production markerAnnotation is entered.
func (s *BaseJava9Listener) EnterMarkerAnnotation(ctx *MarkerAnnotationContext) {}

// ExitMarkerAnnotation is called when production markerAnnotation is exited.
func (s *BaseJava9Listener) ExitMarkerAnnotation(ctx *MarkerAnnotationContext) {}

// EnterSingleElementAnnotation is called when production singleElementAnnotation is entered.
func (s *BaseJava9Listener) EnterSingleElementAnnotation(ctx *SingleElementAnnotationContext) {}

// ExitSingleElementAnnotation is called when production singleElementAnnotation is exited.
func (s *BaseJava9Listener) ExitSingleElementAnnotation(ctx *SingleElementAnnotationContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseJava9Listener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseJava9Listener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterVariableInitializerList is called when production variableInitializerList is entered.
func (s *BaseJava9Listener) EnterVariableInitializerList(ctx *VariableInitializerListContext) {}

// ExitVariableInitializerList is called when production variableInitializerList is exited.
func (s *BaseJava9Listener) ExitVariableInitializerList(ctx *VariableInitializerListContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJava9Listener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJava9Listener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatements is called when production blockStatements is entered.
func (s *BaseJava9Listener) EnterBlockStatements(ctx *BlockStatementsContext) {}

// ExitBlockStatements is called when production blockStatements is exited.
func (s *BaseJava9Listener) ExitBlockStatements(ctx *BlockStatementsContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseJava9Listener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseJava9Listener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is entered.
func (s *BaseJava9Listener) EnterLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// ExitLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is exited.
func (s *BaseJava9Listener) ExitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseJava9Listener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseJava9Listener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJava9Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJava9Listener) ExitStatement(ctx *StatementContext) {}

// EnterStatementNoShortIf is called when production statementNoShortIf is entered.
func (s *BaseJava9Listener) EnterStatementNoShortIf(ctx *StatementNoShortIfContext) {}

// ExitStatementNoShortIf is called when production statementNoShortIf is exited.
func (s *BaseJava9Listener) ExitStatementNoShortIf(ctx *StatementNoShortIfContext) {}

// EnterStatementWithoutTrailingSubstatement is called when production statementWithoutTrailingSubstatement is entered.
func (s *BaseJava9Listener) EnterStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) {
}

// ExitStatementWithoutTrailingSubstatement is called when production statementWithoutTrailingSubstatement is exited.
func (s *BaseJava9Listener) ExitStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) {
}

// EnterEmptyStatement is called when production emptyStatement is entered.
func (s *BaseJava9Listener) EnterEmptyStatement(ctx *EmptyStatementContext) {}

// ExitEmptyStatement is called when production emptyStatement is exited.
func (s *BaseJava9Listener) ExitEmptyStatement(ctx *EmptyStatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *BaseJava9Listener) EnterLabeledStatement(ctx *LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *BaseJava9Listener) ExitLabeledStatement(ctx *LabeledStatementContext) {}

// EnterLabeledStatementNoShortIf is called when production labeledStatementNoShortIf is entered.
func (s *BaseJava9Listener) EnterLabeledStatementNoShortIf(ctx *LabeledStatementNoShortIfContext) {}

// ExitLabeledStatementNoShortIf is called when production labeledStatementNoShortIf is exited.
func (s *BaseJava9Listener) ExitLabeledStatementNoShortIf(ctx *LabeledStatementNoShortIfContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseJava9Listener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseJava9Listener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterStatementExpression is called when production statementExpression is entered.
func (s *BaseJava9Listener) EnterStatementExpression(ctx *StatementExpressionContext) {}

// ExitStatementExpression is called when production statementExpression is exited.
func (s *BaseJava9Listener) ExitStatementExpression(ctx *StatementExpressionContext) {}

// EnterIfThenStatement is called when production ifThenStatement is entered.
func (s *BaseJava9Listener) EnterIfThenStatement(ctx *IfThenStatementContext) {}

// ExitIfThenStatement is called when production ifThenStatement is exited.
func (s *BaseJava9Listener) ExitIfThenStatement(ctx *IfThenStatementContext) {}

// EnterIfThenElseStatement is called when production ifThenElseStatement is entered.
func (s *BaseJava9Listener) EnterIfThenElseStatement(ctx *IfThenElseStatementContext) {}

// ExitIfThenElseStatement is called when production ifThenElseStatement is exited.
func (s *BaseJava9Listener) ExitIfThenElseStatement(ctx *IfThenElseStatementContext) {}

// EnterIfThenElseStatementNoShortIf is called when production ifThenElseStatementNoShortIf is entered.
func (s *BaseJava9Listener) EnterIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) {
}

// ExitIfThenElseStatementNoShortIf is called when production ifThenElseStatementNoShortIf is exited.
func (s *BaseJava9Listener) ExitIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) {
}

// EnterAssertStatement is called when production assertStatement is entered.
func (s *BaseJava9Listener) EnterAssertStatement(ctx *AssertStatementContext) {}

// ExitAssertStatement is called when production assertStatement is exited.
func (s *BaseJava9Listener) ExitAssertStatement(ctx *AssertStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseJava9Listener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseJava9Listener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchBlock is called when production switchBlock is entered.
func (s *BaseJava9Listener) EnterSwitchBlock(ctx *SwitchBlockContext) {}

// ExitSwitchBlock is called when production switchBlock is exited.
func (s *BaseJava9Listener) ExitSwitchBlock(ctx *SwitchBlockContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BaseJava9Listener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BaseJava9Listener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {}

// EnterSwitchLabels is called when production switchLabels is entered.
func (s *BaseJava9Listener) EnterSwitchLabels(ctx *SwitchLabelsContext) {}

// ExitSwitchLabels is called when production switchLabels is exited.
func (s *BaseJava9Listener) ExitSwitchLabels(ctx *SwitchLabelsContext) {}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BaseJava9Listener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BaseJava9Listener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterEnumConstantName is called when production enumConstantName is entered.
func (s *BaseJava9Listener) EnterEnumConstantName(ctx *EnumConstantNameContext) {}

// ExitEnumConstantName is called when production enumConstantName is exited.
func (s *BaseJava9Listener) ExitEnumConstantName(ctx *EnumConstantNameContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseJava9Listener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseJava9Listener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterWhileStatementNoShortIf is called when production whileStatementNoShortIf is entered.
func (s *BaseJava9Listener) EnterWhileStatementNoShortIf(ctx *WhileStatementNoShortIfContext) {}

// ExitWhileStatementNoShortIf is called when production whileStatementNoShortIf is exited.
func (s *BaseJava9Listener) ExitWhileStatementNoShortIf(ctx *WhileStatementNoShortIfContext) {}

// EnterDoStatement is called when production doStatement is entered.
func (s *BaseJava9Listener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production doStatement is exited.
func (s *BaseJava9Listener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseJava9Listener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseJava9Listener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForStatementNoShortIf is called when production forStatementNoShortIf is entered.
func (s *BaseJava9Listener) EnterForStatementNoShortIf(ctx *ForStatementNoShortIfContext) {}

// ExitForStatementNoShortIf is called when production forStatementNoShortIf is exited.
func (s *BaseJava9Listener) ExitForStatementNoShortIf(ctx *ForStatementNoShortIfContext) {}

// EnterBasicForStatement is called when production basicForStatement is entered.
func (s *BaseJava9Listener) EnterBasicForStatement(ctx *BasicForStatementContext) {}

// ExitBasicForStatement is called when production basicForStatement is exited.
func (s *BaseJava9Listener) ExitBasicForStatement(ctx *BasicForStatementContext) {}

// EnterBasicForStatementNoShortIf is called when production basicForStatementNoShortIf is entered.
func (s *BaseJava9Listener) EnterBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) {}

// ExitBasicForStatementNoShortIf is called when production basicForStatementNoShortIf is exited.
func (s *BaseJava9Listener) ExitBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseJava9Listener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseJava9Listener) ExitForInit(ctx *ForInitContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseJava9Listener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseJava9Listener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterStatementExpressionList is called when production statementExpressionList is entered.
func (s *BaseJava9Listener) EnterStatementExpressionList(ctx *StatementExpressionListContext) {}

// ExitStatementExpressionList is called when production statementExpressionList is exited.
func (s *BaseJava9Listener) ExitStatementExpressionList(ctx *StatementExpressionListContext) {}

// EnterEnhancedForStatement is called when production enhancedForStatement is entered.
func (s *BaseJava9Listener) EnterEnhancedForStatement(ctx *EnhancedForStatementContext) {}

// ExitEnhancedForStatement is called when production enhancedForStatement is exited.
func (s *BaseJava9Listener) ExitEnhancedForStatement(ctx *EnhancedForStatementContext) {}

// EnterEnhancedForStatementNoShortIf is called when production enhancedForStatementNoShortIf is entered.
func (s *BaseJava9Listener) EnterEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) {
}

// ExitEnhancedForStatementNoShortIf is called when production enhancedForStatementNoShortIf is exited.
func (s *BaseJava9Listener) ExitEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) {
}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseJava9Listener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseJava9Listener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseJava9Listener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseJava9Listener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseJava9Listener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseJava9Listener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseJava9Listener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseJava9Listener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterSynchronizedStatement is called when production synchronizedStatement is entered.
func (s *BaseJava9Listener) EnterSynchronizedStatement(ctx *SynchronizedStatementContext) {}

// ExitSynchronizedStatement is called when production synchronizedStatement is exited.
func (s *BaseJava9Listener) ExitSynchronizedStatement(ctx *SynchronizedStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseJava9Listener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseJava9Listener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterCatches is called when production catches is entered.
func (s *BaseJava9Listener) EnterCatches(ctx *CatchesContext) {}

// ExitCatches is called when production catches is exited.
func (s *BaseJava9Listener) ExitCatches(ctx *CatchesContext) {}

// EnterCatchClause is called when production catchClause is entered.
func (s *BaseJava9Listener) EnterCatchClause(ctx *CatchClauseContext) {}

// ExitCatchClause is called when production catchClause is exited.
func (s *BaseJava9Listener) ExitCatchClause(ctx *CatchClauseContext) {}

// EnterCatchFormalParameter is called when production catchFormalParameter is entered.
func (s *BaseJava9Listener) EnterCatchFormalParameter(ctx *CatchFormalParameterContext) {}

// ExitCatchFormalParameter is called when production catchFormalParameter is exited.
func (s *BaseJava9Listener) ExitCatchFormalParameter(ctx *CatchFormalParameterContext) {}

// EnterCatchType is called when production catchType is entered.
func (s *BaseJava9Listener) EnterCatchType(ctx *CatchTypeContext) {}

// ExitCatchType is called when production catchType is exited.
func (s *BaseJava9Listener) ExitCatchType(ctx *CatchTypeContext) {}

// EnterFinally_ is called when production finally_ is entered.
func (s *BaseJava9Listener) EnterFinally_(ctx *Finally_Context) {}

// ExitFinally_ is called when production finally_ is exited.
func (s *BaseJava9Listener) ExitFinally_(ctx *Finally_Context) {}

// EnterTryWithResourcesStatement is called when production tryWithResourcesStatement is entered.
func (s *BaseJava9Listener) EnterTryWithResourcesStatement(ctx *TryWithResourcesStatementContext) {}

// ExitTryWithResourcesStatement is called when production tryWithResourcesStatement is exited.
func (s *BaseJava9Listener) ExitTryWithResourcesStatement(ctx *TryWithResourcesStatementContext) {}

// EnterResourceSpecification is called when production resourceSpecification is entered.
func (s *BaseJava9Listener) EnterResourceSpecification(ctx *ResourceSpecificationContext) {}

// ExitResourceSpecification is called when production resourceSpecification is exited.
func (s *BaseJava9Listener) ExitResourceSpecification(ctx *ResourceSpecificationContext) {}

// EnterResourceList is called when production resourceList is entered.
func (s *BaseJava9Listener) EnterResourceList(ctx *ResourceListContext) {}

// ExitResourceList is called when production resourceList is exited.
func (s *BaseJava9Listener) ExitResourceList(ctx *ResourceListContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseJava9Listener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseJava9Listener) ExitResource(ctx *ResourceContext) {}

// EnterVariableAccess is called when production variableAccess is entered.
func (s *BaseJava9Listener) EnterVariableAccess(ctx *VariableAccessContext) {}

// ExitVariableAccess is called when production variableAccess is exited.
func (s *BaseJava9Listener) ExitVariableAccess(ctx *VariableAccessContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseJava9Listener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseJava9Listener) ExitPrimary(ctx *PrimaryContext) {}

// EnterPrimaryNoNewArray is called when production primaryNoNewArray is entered.
func (s *BaseJava9Listener) EnterPrimaryNoNewArray(ctx *PrimaryNoNewArrayContext) {}

// ExitPrimaryNoNewArray is called when production primaryNoNewArray is exited.
func (s *BaseJava9Listener) ExitPrimaryNoNewArray(ctx *PrimaryNoNewArrayContext) {}

// EnterPrimaryNoNewArray_lf_arrayAccess is called when production primaryNoNewArray_lf_arrayAccess is entered.
func (s *BaseJava9Listener) EnterPrimaryNoNewArray_lf_arrayAccess(ctx *PrimaryNoNewArray_lf_arrayAccessContext) {
}

// ExitPrimaryNoNewArray_lf_arrayAccess is called when production primaryNoNewArray_lf_arrayAccess is exited.
func (s *BaseJava9Listener) ExitPrimaryNoNewArray_lf_arrayAccess(ctx *PrimaryNoNewArray_lf_arrayAccessContext) {
}

// EnterPrimaryNoNewArray_lfno_arrayAccess is called when production primaryNoNewArray_lfno_arrayAccess is entered.
func (s *BaseJava9Listener) EnterPrimaryNoNewArray_lfno_arrayAccess(ctx *PrimaryNoNewArray_lfno_arrayAccessContext) {
}

// ExitPrimaryNoNewArray_lfno_arrayAccess is called when production primaryNoNewArray_lfno_arrayAccess is exited.
func (s *BaseJava9Listener) ExitPrimaryNoNewArray_lfno_arrayAccess(ctx *PrimaryNoNewArray_lfno_arrayAccessContext) {
}

// EnterPrimaryNoNewArray_lf_primary is called when production primaryNoNewArray_lf_primary is entered.
func (s *BaseJava9Listener) EnterPrimaryNoNewArray_lf_primary(ctx *PrimaryNoNewArray_lf_primaryContext) {
}

// ExitPrimaryNoNewArray_lf_primary is called when production primaryNoNewArray_lf_primary is exited.
func (s *BaseJava9Listener) ExitPrimaryNoNewArray_lf_primary(ctx *PrimaryNoNewArray_lf_primaryContext) {
}

// EnterPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is called when production primaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is entered.
func (s *BaseJava9Listener) EnterPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primaryContext) {
}

// ExitPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is called when production primaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is exited.
func (s *BaseJava9Listener) ExitPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primaryContext) {
}

// EnterPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is called when production primaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is entered.
func (s *BaseJava9Listener) EnterPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primaryContext) {
}

// ExitPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is called when production primaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is exited.
func (s *BaseJava9Listener) ExitPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primaryContext) {
}

// EnterPrimaryNoNewArray_lfno_primary is called when production primaryNoNewArray_lfno_primary is entered.
func (s *BaseJava9Listener) EnterPrimaryNoNewArray_lfno_primary(ctx *PrimaryNoNewArray_lfno_primaryContext) {
}

// ExitPrimaryNoNewArray_lfno_primary is called when production primaryNoNewArray_lfno_primary is exited.
func (s *BaseJava9Listener) ExitPrimaryNoNewArray_lfno_primary(ctx *PrimaryNoNewArray_lfno_primaryContext) {
}

// EnterPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is called when production primaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is entered.
func (s *BaseJava9Listener) EnterPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primaryContext) {
}

// ExitPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is called when production primaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is exited.
func (s *BaseJava9Listener) ExitPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primaryContext) {
}

// EnterPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is called when production primaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is entered.
func (s *BaseJava9Listener) EnterPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primaryContext) {
}

// ExitPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is called when production primaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is exited.
func (s *BaseJava9Listener) ExitPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primaryContext) {
}

// EnterClassLiteral is called when production classLiteral is entered.
func (s *BaseJava9Listener) EnterClassLiteral(ctx *ClassLiteralContext) {}

// ExitClassLiteral is called when production classLiteral is exited.
func (s *BaseJava9Listener) ExitClassLiteral(ctx *ClassLiteralContext) {}

// EnterClassInstanceCreationExpression is called when production classInstanceCreationExpression is entered.
func (s *BaseJava9Listener) EnterClassInstanceCreationExpression(ctx *ClassInstanceCreationExpressionContext) {
}

// ExitClassInstanceCreationExpression is called when production classInstanceCreationExpression is exited.
func (s *BaseJava9Listener) ExitClassInstanceCreationExpression(ctx *ClassInstanceCreationExpressionContext) {
}

// EnterClassInstanceCreationExpression_lf_primary is called when production classInstanceCreationExpression_lf_primary is entered.
func (s *BaseJava9Listener) EnterClassInstanceCreationExpression_lf_primary(ctx *ClassInstanceCreationExpression_lf_primaryContext) {
}

// ExitClassInstanceCreationExpression_lf_primary is called when production classInstanceCreationExpression_lf_primary is exited.
func (s *BaseJava9Listener) ExitClassInstanceCreationExpression_lf_primary(ctx *ClassInstanceCreationExpression_lf_primaryContext) {
}

// EnterClassInstanceCreationExpression_lfno_primary is called when production classInstanceCreationExpression_lfno_primary is entered.
func (s *BaseJava9Listener) EnterClassInstanceCreationExpression_lfno_primary(ctx *ClassInstanceCreationExpression_lfno_primaryContext) {
}

// ExitClassInstanceCreationExpression_lfno_primary is called when production classInstanceCreationExpression_lfno_primary is exited.
func (s *BaseJava9Listener) ExitClassInstanceCreationExpression_lfno_primary(ctx *ClassInstanceCreationExpression_lfno_primaryContext) {
}

// EnterTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is entered.
func (s *BaseJava9Listener) EnterTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// ExitTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is exited.
func (s *BaseJava9Listener) ExitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// EnterFieldAccess is called when production fieldAccess is entered.
func (s *BaseJava9Listener) EnterFieldAccess(ctx *FieldAccessContext) {}

// ExitFieldAccess is called when production fieldAccess is exited.
func (s *BaseJava9Listener) ExitFieldAccess(ctx *FieldAccessContext) {}

// EnterFieldAccess_lf_primary is called when production fieldAccess_lf_primary is entered.
func (s *BaseJava9Listener) EnterFieldAccess_lf_primary(ctx *FieldAccess_lf_primaryContext) {}

// ExitFieldAccess_lf_primary is called when production fieldAccess_lf_primary is exited.
func (s *BaseJava9Listener) ExitFieldAccess_lf_primary(ctx *FieldAccess_lf_primaryContext) {}

// EnterFieldAccess_lfno_primary is called when production fieldAccess_lfno_primary is entered.
func (s *BaseJava9Listener) EnterFieldAccess_lfno_primary(ctx *FieldAccess_lfno_primaryContext) {}

// ExitFieldAccess_lfno_primary is called when production fieldAccess_lfno_primary is exited.
func (s *BaseJava9Listener) ExitFieldAccess_lfno_primary(ctx *FieldAccess_lfno_primaryContext) {}

// EnterArrayAccess is called when production arrayAccess is entered.
func (s *BaseJava9Listener) EnterArrayAccess(ctx *ArrayAccessContext) {}

// ExitArrayAccess is called when production arrayAccess is exited.
func (s *BaseJava9Listener) ExitArrayAccess(ctx *ArrayAccessContext) {}

// EnterArrayAccess_lf_primary is called when production arrayAccess_lf_primary is entered.
func (s *BaseJava9Listener) EnterArrayAccess_lf_primary(ctx *ArrayAccess_lf_primaryContext) {}

// ExitArrayAccess_lf_primary is called when production arrayAccess_lf_primary is exited.
func (s *BaseJava9Listener) ExitArrayAccess_lf_primary(ctx *ArrayAccess_lf_primaryContext) {}

// EnterArrayAccess_lfno_primary is called when production arrayAccess_lfno_primary is entered.
func (s *BaseJava9Listener) EnterArrayAccess_lfno_primary(ctx *ArrayAccess_lfno_primaryContext) {}

// ExitArrayAccess_lfno_primary is called when production arrayAccess_lfno_primary is exited.
func (s *BaseJava9Listener) ExitArrayAccess_lfno_primary(ctx *ArrayAccess_lfno_primaryContext) {}

// EnterMethodInvocation is called when production methodInvocation is entered.
func (s *BaseJava9Listener) EnterMethodInvocation(ctx *MethodInvocationContext) {}

// ExitMethodInvocation is called when production methodInvocation is exited.
func (s *BaseJava9Listener) ExitMethodInvocation(ctx *MethodInvocationContext) {}

// EnterMethodInvocation_lf_primary is called when production methodInvocation_lf_primary is entered.
func (s *BaseJava9Listener) EnterMethodInvocation_lf_primary(ctx *MethodInvocation_lf_primaryContext) {
}

// ExitMethodInvocation_lf_primary is called when production methodInvocation_lf_primary is exited.
func (s *BaseJava9Listener) ExitMethodInvocation_lf_primary(ctx *MethodInvocation_lf_primaryContext) {}

// EnterMethodInvocation_lfno_primary is called when production methodInvocation_lfno_primary is entered.
func (s *BaseJava9Listener) EnterMethodInvocation_lfno_primary(ctx *MethodInvocation_lfno_primaryContext) {
}

// ExitMethodInvocation_lfno_primary is called when production methodInvocation_lfno_primary is exited.
func (s *BaseJava9Listener) ExitMethodInvocation_lfno_primary(ctx *MethodInvocation_lfno_primaryContext) {
}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseJava9Listener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseJava9Listener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterMethodReference is called when production methodReference is entered.
func (s *BaseJava9Listener) EnterMethodReference(ctx *MethodReferenceContext) {}

// ExitMethodReference is called when production methodReference is exited.
func (s *BaseJava9Listener) ExitMethodReference(ctx *MethodReferenceContext) {}

// EnterMethodReference_lf_primary is called when production methodReference_lf_primary is entered.
func (s *BaseJava9Listener) EnterMethodReference_lf_primary(ctx *MethodReference_lf_primaryContext) {}

// ExitMethodReference_lf_primary is called when production methodReference_lf_primary is exited.
func (s *BaseJava9Listener) ExitMethodReference_lf_primary(ctx *MethodReference_lf_primaryContext) {}

// EnterMethodReference_lfno_primary is called when production methodReference_lfno_primary is entered.
func (s *BaseJava9Listener) EnterMethodReference_lfno_primary(ctx *MethodReference_lfno_primaryContext) {
}

// ExitMethodReference_lfno_primary is called when production methodReference_lfno_primary is exited.
func (s *BaseJava9Listener) ExitMethodReference_lfno_primary(ctx *MethodReference_lfno_primaryContext) {
}

// EnterArrayCreationExpression is called when production arrayCreationExpression is entered.
func (s *BaseJava9Listener) EnterArrayCreationExpression(ctx *ArrayCreationExpressionContext) {}

// ExitArrayCreationExpression is called when production arrayCreationExpression is exited.
func (s *BaseJava9Listener) ExitArrayCreationExpression(ctx *ArrayCreationExpressionContext) {}

// EnterDimExprs is called when production dimExprs is entered.
func (s *BaseJava9Listener) EnterDimExprs(ctx *DimExprsContext) {}

// ExitDimExprs is called when production dimExprs is exited.
func (s *BaseJava9Listener) ExitDimExprs(ctx *DimExprsContext) {}

// EnterDimExpr is called when production dimExpr is entered.
func (s *BaseJava9Listener) EnterDimExpr(ctx *DimExprContext) {}

// ExitDimExpr is called when production dimExpr is exited.
func (s *BaseJava9Listener) ExitDimExpr(ctx *DimExprContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BaseJava9Listener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BaseJava9Listener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJava9Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJava9Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *BaseJava9Listener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *BaseJava9Listener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterLambdaParameters is called when production lambdaParameters is entered.
func (s *BaseJava9Listener) EnterLambdaParameters(ctx *LambdaParametersContext) {}

// ExitLambdaParameters is called when production lambdaParameters is exited.
func (s *BaseJava9Listener) ExitLambdaParameters(ctx *LambdaParametersContext) {}

// EnterInferredFormalParameterList is called when production inferredFormalParameterList is entered.
func (s *BaseJava9Listener) EnterInferredFormalParameterList(ctx *InferredFormalParameterListContext) {
}

// ExitInferredFormalParameterList is called when production inferredFormalParameterList is exited.
func (s *BaseJava9Listener) ExitInferredFormalParameterList(ctx *InferredFormalParameterListContext) {}

// EnterLambdaBody is called when production lambdaBody is entered.
func (s *BaseJava9Listener) EnterLambdaBody(ctx *LambdaBodyContext) {}

// ExitLambdaBody is called when production lambdaBody is exited.
func (s *BaseJava9Listener) ExitLambdaBody(ctx *LambdaBodyContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseJava9Listener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseJava9Listener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseJava9Listener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseJava9Listener) ExitAssignment(ctx *AssignmentContext) {}

// EnterLeftHandSide is called when production leftHandSide is entered.
func (s *BaseJava9Listener) EnterLeftHandSide(ctx *LeftHandSideContext) {}

// ExitLeftHandSide is called when production leftHandSide is exited.
func (s *BaseJava9Listener) ExitLeftHandSide(ctx *LeftHandSideContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseJava9Listener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseJava9Listener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseJava9Listener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseJava9Listener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterConditionalOrExpression is called when production conditionalOrExpression is entered.
func (s *BaseJava9Listener) EnterConditionalOrExpression(ctx *ConditionalOrExpressionContext) {}

// ExitConditionalOrExpression is called when production conditionalOrExpression is exited.
func (s *BaseJava9Listener) ExitConditionalOrExpression(ctx *ConditionalOrExpressionContext) {}

// EnterConditionalAndExpression is called when production conditionalAndExpression is entered.
func (s *BaseJava9Listener) EnterConditionalAndExpression(ctx *ConditionalAndExpressionContext) {}

// ExitConditionalAndExpression is called when production conditionalAndExpression is exited.
func (s *BaseJava9Listener) ExitConditionalAndExpression(ctx *ConditionalAndExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *BaseJava9Listener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *BaseJava9Listener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *BaseJava9Listener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *BaseJava9Listener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseJava9Listener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseJava9Listener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseJava9Listener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseJava9Listener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseJava9Listener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseJava9Listener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseJava9Listener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseJava9Listener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseJava9Listener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseJava9Listener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseJava9Listener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseJava9Listener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseJava9Listener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseJava9Listener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPreIncrementExpression is called when production preIncrementExpression is entered.
func (s *BaseJava9Listener) EnterPreIncrementExpression(ctx *PreIncrementExpressionContext) {}

// ExitPreIncrementExpression is called when production preIncrementExpression is exited.
func (s *BaseJava9Listener) ExitPreIncrementExpression(ctx *PreIncrementExpressionContext) {}

// EnterPreDecrementExpression is called when production preDecrementExpression is entered.
func (s *BaseJava9Listener) EnterPreDecrementExpression(ctx *PreDecrementExpressionContext) {}

// ExitPreDecrementExpression is called when production preDecrementExpression is exited.
func (s *BaseJava9Listener) ExitPreDecrementExpression(ctx *PreDecrementExpressionContext) {}

// EnterUnaryExpressionNotPlusMinus is called when production unaryExpressionNotPlusMinus is entered.
func (s *BaseJava9Listener) EnterUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) {
}

// ExitUnaryExpressionNotPlusMinus is called when production unaryExpressionNotPlusMinus is exited.
func (s *BaseJava9Listener) ExitUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseJava9Listener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseJava9Listener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterPostIncrementExpression is called when production postIncrementExpression is entered.
func (s *BaseJava9Listener) EnterPostIncrementExpression(ctx *PostIncrementExpressionContext) {}

// ExitPostIncrementExpression is called when production postIncrementExpression is exited.
func (s *BaseJava9Listener) ExitPostIncrementExpression(ctx *PostIncrementExpressionContext) {}

// EnterPostIncrementExpression_lf_postfixExpression is called when production postIncrementExpression_lf_postfixExpression is entered.
func (s *BaseJava9Listener) EnterPostIncrementExpression_lf_postfixExpression(ctx *PostIncrementExpression_lf_postfixExpressionContext) {
}

// ExitPostIncrementExpression_lf_postfixExpression is called when production postIncrementExpression_lf_postfixExpression is exited.
func (s *BaseJava9Listener) ExitPostIncrementExpression_lf_postfixExpression(ctx *PostIncrementExpression_lf_postfixExpressionContext) {
}

// EnterPostDecrementExpression is called when production postDecrementExpression is entered.
func (s *BaseJava9Listener) EnterPostDecrementExpression(ctx *PostDecrementExpressionContext) {}

// ExitPostDecrementExpression is called when production postDecrementExpression is exited.
func (s *BaseJava9Listener) ExitPostDecrementExpression(ctx *PostDecrementExpressionContext) {}

// EnterPostDecrementExpression_lf_postfixExpression is called when production postDecrementExpression_lf_postfixExpression is entered.
func (s *BaseJava9Listener) EnterPostDecrementExpression_lf_postfixExpression(ctx *PostDecrementExpression_lf_postfixExpressionContext) {
}

// ExitPostDecrementExpression_lf_postfixExpression is called when production postDecrementExpression_lf_postfixExpression is exited.
func (s *BaseJava9Listener) ExitPostDecrementExpression_lf_postfixExpression(ctx *PostDecrementExpression_lf_postfixExpressionContext) {
}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseJava9Listener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseJava9Listener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseJava9Listener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseJava9Listener) ExitIdentifier(ctx *IdentifierContext) {}
