<?php

namespace App\Http\Requests;

class ClassSubjectMaterialCheckRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'class_subject_material_id' => 'required|numeric|gt:0',
            'student_id' => 'required|numeric|gt:0',
        ];
    }

    /**
     * @return array
     * Custom validation message
     */
    public function messages(): array
    {
        return [
            'class_subject_material_id.required' => 'The class subject material ID field is required.',
            'class_subject_material_id.numeric' => 'The class subject material ID must be a number.',
            'class_subject_id.gt' => 'The class subject material ID must be a positive number.',
            'student_id.required' => 'The class student ID field is required.',
            'student_id.numeric' => 'The class student ID must be a number.',
            'student_id.gt' => 'The class student ID must be a positive number.',
        ];
    }
}
