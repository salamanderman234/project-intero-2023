<?php

namespace App\Http\Requests;

class ClassSubjectAssignmentSubmissionRequest extends FormRequest
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
            'class_subject_assignment_id' => 'required|numeric|gt:0',
            'student_id' => 'required|numeric|gt:0',
            'attachment' => 'required',
        ];
    }

    /**
     * @return array
     * Custom validation message
     */
    public function messages(): array
    {
        return [
            'class_subject_assignment_id.required' => 'The class subject assignment ID field is required.',
            'class_subject_assignment_id.numeric' => 'The class subject assignment ID must be a number.',
            'class_subject_assignment_id.gt' => 'The class subject assignment ID must be a positive number.',
            'student_id.required' => 'The class student ID field is required.',
            'student_id.numeric' => 'The class student ID must be a number.',
            'student_id.gt' => 'The class student ID must be a positive number.',
            'attachment.required' => 'The attachment field is required.',
        ];
    }
}
